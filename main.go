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
	for i, p := range players {
		ratedPlayers[i] = RatePosition(p, position)
	}
	Sort(*sortFlag, ratedPlayers)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredDark)

	header := makeHeader(position)
	t.AppendHeader(header)
	for _, rp := range ratedPlayers {
		t.AppendRow(makeRow(rp))
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

func makeRow(rp *RatedPlayer) []interface{} {
	row := []interface{}{rp.player.name, rp.player.age, rp.player.position, fmt.Sprintf("%.3f", rp.averageRating)}
	for _, v := range rp.ratings {
		row = append(row, fmt.Sprintf("%.3f", v))
	}
	return row
}
