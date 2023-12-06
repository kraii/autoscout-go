package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"sort"
)

type ByRating struct {
	pos string
	rp  []*RatedPlayer
}

func (a ByRating) Len() int           { return len(a.rp) }
func (a ByRating) Less(i, j int) bool { return a.rp[i].ratings[a.pos] < a.rp[j].ratings[a.pos] }
func (a ByRating) Swap(i, j int)      { a.rp[i], a.rp[j] = a.rp[j], a.rp[i] }

func main() {
	players, err := Parse("/media/kraii/Windows/Users/matth/Documents/scout/dm.html")
	if err != nil {
		panic(err)
	}

	ratedPlayers := make([]*RatedPlayer, len(players))
	position := Positions["DM"]

	for i, p := range players {
		ratedPlayers[i] = RatePosition(p, position)
	}
	sort.Sort(ByRating{"DLP-S", ratedPlayers})

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

func makeHeader(position Position) []interface{} {
	header := []interface{}{"Name", "Age", "Position", "Average"}
	for roleName, role := range position {
		header = append(header, fmt.Sprintf("%s-%s", roleName, role.duty))
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
