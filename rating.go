package main

type RoleAttributes struct {
	primaryAttributes, secondaryAttributes []func(*Player) int
}

var DefensiveMid = RoleAttributes{
	primaryAttributes:   []func(*Player) int{tackling, anticipation, concentration, positioning, teamwork},
	secondaryAttributes: []func(*Player) int{firstTouch, marking, passing, aggression, composure, decisions, workRate, stamina, strength},
}

const maxAttributeValue = 20.0

func Rate(p *Player, r *RoleAttributes) float64 {
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
