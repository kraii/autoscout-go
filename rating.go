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
	"GK": {
		Role{
			name:                "GK",
			duty:                "D",
			primaryAttributes:   []func(*Player) int{aerialReach, commandOfArea, communication, handling, kicking, reflexes, concentration, positioning, agility},
			secondaryAttributes: []func(*Player) int{oneOnOnes, throwing, anticipation, decisions},
		},
	},
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
	"W": {
		Role{
			name:                "IW",
			duty:                "S",
			primaryAttributes:   []func(*Player) int{crossing, dribbling, passing, technique, acceleration, agility},
			secondaryAttributes: []func(*Player) int{firstTouch, longShots, composure, decisions, offTheBall, vision, workRate, balance, passing, stamina},
		},
	},
	"ST": {
		Role{
			name:                "AF",
			duty:                "A",
			primaryAttributes:   []func(*Player) int{finishing, dribbling, firstTouch, technique, composure, offTheBall, acceleration},
			secondaryAttributes: []func(*Player) int{passing, decisions, anticipation, workRate, agility, balance, pace, stamina},
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

func acceleration(p *Player) int  { return p.acceleration }
func aerialReach(p *Player) int   { return p.aerialReach }
func aggression(p *Player) int    { return p.aggression }
func agility(p *Player) int       { return p.agility }
func anticipation(p *Player) int  { return p.anticipation }
func balance(p *Player) int       { return p.balance }
func bravery(p *Player) int       { return p.bravery }
func commandOfArea(p *Player) int { return p.commandOfArea }
func communication(p *Player) int { return p.communication }
func composure(p *Player) int     { return p.composure }
func concentration(p *Player) int { return p.concentration }
func crossing(p *Player) int      { return p.crossing }
func decisions(p *Player) int     { return p.decisions }
func dribbling(p *Player) int     { return p.dribbling }
func finishing(p *Player) int     { return p.finishing }
func firstTouch(p *Player) int    { return p.firstTouch }
func handling(p *Player) int      { return p.handling }
func kicking(p *Player) int       { return p.kicking }
func longShots(p *Player) int     { return p.longShots }
func marking(p *Player) int       { return p.marking }
func offTheBall(p *Player) int    { return p.offTheBall }
func oneOnOnes(p *Player) int     { return p.oneOnOnes }
func pace(p *Player) int          { return p.pace }
func passing(p *Player) int       { return p.passing }
func positioning(p *Player) int   { return p.positioning }
func reflexes(p *Player) int      { return p.reflexes }
func stamina(p *Player) int       { return p.stamina }
func strength(p *Player) int      { return p.strength }
func tackling(p *Player) int      { return p.tackling }
func teamwork(p *Player) int      { return p.teamwork }
func technique(p *Player) int     { return p.technique }
func throwing(p *Player) int      { return p.throwing }
func vision(p *Player) int        { return p.vision }
func workRate(p *Player) int      { return p.workRate }
