package main

import "fmt"

func main() {
	players, err := Parse("/media/kraii/Windows/Users/matth/Documents/scout/team.html")
	if err != nil {
		panic(err)
	}

	for _, p := range players {
		fmt.Printf("%s rated %.3f\n", p.name, Rate(p, &DefensiveMid))
	}
}
