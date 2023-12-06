package main

import (
	"sort"
)

const Average = "average"

type ByRating struct {
	pos string
	rp  []*RatedPlayer
}

func (a ByRating) Len() int           { return len(a.rp) }
func (a ByRating) Less(i, j int) bool { return a.rp[i].ratings[a.pos] < a.rp[j].ratings[a.pos] }
func (a ByRating) Swap(i, j int)      { a.rp[i], a.rp[j] = a.rp[j], a.rp[i] }

type ByAverage []*RatedPlayer

func (a ByAverage) Len() int           { return len(a) }
func (a ByAverage) Less(i, j int) bool { return a[i].averageRating < a[j].averageRating }
func (a ByAverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
	var sorter sort.Interface
	if sortArg == Average {
		sorter = ByAverage(players)
	} else {
		sorter = ByRating{sortArg, players}
	}

	sort.Sort(sorter)
}
