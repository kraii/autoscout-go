package main

import (
	"fmt"
	"sort"
)

type RatedPlayer struct {
	player *Player
	rating float64
}

type ByRating []*RatedPlayer

func (a ByRating) Len() int           { return len(a) }
func (a ByRating) Less(i, j int) bool { return a[i].rating < a[j].rating }
func (a ByRating) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	players, err := Parse("/media/kraii/Windows/Users/matth/Documents/scout/dm.html")
	if err != nil {
		panic(err)
	}

	ratedPlayers := make([]*RatedPlayer, len(players))
	for i, p := range players {
		ratedPlayers[i] = &RatedPlayer{
			p,
			Rate(p, &DefensiveMid),
		}
	}
	sort.Sort(ByRating(ratedPlayers))

	for _, rp := range ratedPlayers {
		fmt.Printf("%s rated %.3f\n", rp.player.name, rp.rating)
	}
}
