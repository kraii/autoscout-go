package main

import (
	"flag"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func main() {
	var fileFlag = flag.String("f", "-1", "the file html file to read from, exported from FM")
	var positionFlag = flag.String("p", "-1", "the position to rate for, e.g. DM, WB, CB, W, ST")
	var sortFlag = flag.String("s", Average, "Which role to sort by, e.g. DLP-S, AF-A")
	var maxAgeFlag = flag.Int("a", 99, "Max age of player")
	flag.Parse()
	checkFlags(fileFlag, positionFlag)

	position, posExisted := Positions[*positionFlag]
	if !posExisted {
		println("Unknown position", *positionFlag)
		os.Exit(2)
	}
	if !CheckSortArg(*sortFlag, position) {
		println("Invalid sort flag:", *sortFlag)
		os.Exit(3)
	}

	players, err := Parse(*fileFlag)

	if err != nil {
		println(err.Error())
		os.Exit(4)
	}

	ratedPlayers := make([]*RatedPlayer, len(players))
	i := 0
	for _, p := range players {
		if p.age <= *maxAgeFlag {
			ratedPlayers[i] = RatePosition(p, position)
			i++
		}
	}
	ratedPlayers = ratedPlayers[:i]

	Sort(*sortFlag, ratedPlayers)

	renderTable(position, ratedPlayers)
}

func renderTable(position Position, ratedPlayers []*RatedPlayer) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredDark)

	header := makeHeader(position)
	t.AppendHeader(header)
	for _, rp := range ratedPlayers {
		t.AppendRow(makeRow(position, rp))
	}
	t.Render()
}

func checkFlags(flags ...*string) {
	for _, f := range flags {
		if *f == "-1" {
			flag.Usage()
			os.Exit(1)
		}
	}
}

func makeHeader(position Position) []interface{} {
	header := []interface{}{"Name", "Age", "Position", "Average"}
	for _, role := range position {
		header = append(header, role.Format())
	}

	return header
}

func makeRow(position Position, rp *RatedPlayer) []interface{} {
	row := []interface{}{rp.player.name, rp.player.age, rp.player.position, fmt.Sprintf("%.3f", rp.averageRating)}
	for _, role := range position {
		rating := rp.ratings[role.Format()]
		row = append(row, fmt.Sprintf("%.3f", rating))
	}
	return row
}
