package main

import "fmt"

type Role struct {
	name                                   string
	duty                                   string
	primaryAttributes, secondaryAttributes []func(*Player) int
}

func (r *Role) Format() string {
	return fmt.Sprintf("%s-%s", r.name, r.duty)
}

type Position = []Role

var Positions = map[string]Position{
	"DM": {
		Role{
			name:                "DM",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{tackling, anticipation, concentration, positioning, teamwork},
			secondaryAttributes: []func(*Player) int{firstTouch, marking, passing, aggression, composure, decisions, workRate, stamina, strength},
		},
		Role{
			name:                "DLP",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{firstTouch, passing, technique, composure, decisions, teamwork, vision},
			secondaryAttributes: []func(*Player) int{anticipation, offTheBall, positioning},
		},
		Role{
			name:                "BWM",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{tackling, aggression, anticipation, teamwork, workRate, stamina},
			secondaryAttributes: []func(*Player) int{marking, passing, bravery, agility, pace, strength},
		},
	},
	"CM": {
		Role{
			name:                "MEZ",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{passing, technique, decisions, offTheBall, workRate, acceleration},
			secondaryAttributes: []func(*Player) int{dribbling, firstTouch, longShots, tackling, anticipation, composure, vision, balance, stamina},
		},
		Role{
			name:                "DLP",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{firstTouch, passing, technique, communication, decisions, teamwork, vision},
			secondaryAttributes: []func(*Player) int{anticipation, offTheBall, positioning, balance},
		},
		Role{
			name:                "CM",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{firstTouch, passing, tackling, decisions, teamwork},
			secondaryAttributes: []func(*Player) int{technique, anticipation, composure, concentration, offTheBall, vision, workRate, stamina},
		},
	},
}

type RatedPlayer struct {
	player        *Player
	averageRating float64
	ratings       map[string]float64
}

func RatePosition(p *Player, position Position) *RatedPlayer {
	ratings := make(map[string]float64)
	total := 0.0
	for _, role := range position {
		rating := Rate(p, &role)
		ratings[role.Format()] = rating
		total += rating
	}
	return &RatedPlayer{
		player:        p,
		averageRating: total / float64(len(position)),
		ratings:       ratings,
	}
}

const maxAttributeValue = 20.0

func Rate(p *Player, r *Role) float64 {
	const primaryFactor = 1.5
	const secondaryFactor = 1.0

	total := 0.0

	for _, attribute := range r.primaryAttributes {
		total += primaryFactor * float64(attribute(p))
	}

	for _, attribute := range r.secondaryAttributes {
		total += secondaryFactor * float64(attribute(p))
	}

	maximum := float64(len(r.primaryAttributes)) * primaryFactor * maxAttributeValue
	maximum += float64(len(r.secondaryAttributes)) * secondaryFactor * maxAttributeValue
	return (total / maximum) * 20
}

func tackling(p *Player) int      { return p.tackling }
func anticipation(p *Player) int  { return p.anticipation }
func concentration(p *Player) int { return p.concentration }
func positioning(p *Player) int   { return p.positioning }
func teamwork(p *Player) int      { return p.teamwork }
func firstTouch(p *Player) int    { return p.firstTouch }
func marking(p *Player) int       { return p.marking }
func passing(p *Player) int       { return p.passing }
func aggression(p *Player) int    { return p.aggression }
func composure(p *Player) int     { return p.composure }
func decisions(p *Player) int     { return p.decisions }
func workRate(p *Player) int      { return p.workRate }
func stamina(p *Player) int       { return p.stamina }
func strength(p *Player) int      { return p.strength }
func technique(p *Player) int     { return p.technique }
func vision(p *Player) int        { return p.vision }
func offTheBall(p *Player) int    { return p.offTheBall }
func bravery(p *Player) int       { return p.bravery }
func agility(p *Player) int       { return p.agility }
func pace(p *Player) int          { return p.pace }
func acceleration(p *Player) int  { return p.acceleration }
func dribbling(p *Player) int     { return p.dribbling }
func longShots(p *Player) int     { return p.longShots }
func communication(p *Player) int { return p.communication }
func balance(p *Player) int       { return p.balance }
