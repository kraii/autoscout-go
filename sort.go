package main

import (
	"cmp"
	"log"
	"slices"
)

const Average = "average"

func CheckSortArg(sortArg string, position Position) bool {
	if sortArg == Average {
		return true
	}
	for _, r := range position {
		if sortArg == r.Format() {
			return true
		}
	}
	return false
}

func Sort(sortArg string, players []*RatedPlayer) {
	var sortValue func(p *RatedPlayer) float64
	if sortArg == Average {
		sortValue = func(p *RatedPlayer) float64 { return p.averageRating }
	} else {
		sortValue = func(p *RatedPlayer) float64 {
			f := p.ratings[sortArg]
			log.Printf("%s - %.3f", p.player.name, f)
			return f
		}
	}

	slices.SortFunc(players, func(a, b *RatedPlayer) int {
		return cmp.Compare(sortValue(a), sortValue(b))
	})
}
