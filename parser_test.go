package main

import (
	testify "github.com/stretchr/testify/assert"
	"path"
	"runtime"
	"testing"
)

func TestParseFile(t *testing.T) {
	assert := testify.New(t)
	_, filename, _, _ := runtime.Caller(0)
	players, err := Parse(path.Join(path.Dir(filename), "test.html"))

	assert.Nil(err)
	assert.Len(players, 2)

	tomDonaghy := players[0]
	harryBrumwell := players[1]

	assert.Equal(&expectedTom, tomDonaghy)
	assert.Equal(&expectedHarry, harryBrumwell)
}

var expectedTom = Player{
	name:           "Tom Donaghy",
	age:            20,
	position:       "GK",
	corners:        3,
	crossing:       3,
	dribbling:      2,
	finishing:      2,
	firstTouch:     1,
	freeKickTaking: 4,
	heading:        1,
	longShots:      2,
	throwIns:       3,
	marking:        2,
	passing:        8,
	penaltyTaking:  3,
	tackling:       2,
	technique:      7,
	acceleration:   5,
	agility:        6,
	balance:        8,
	jumping:        10,
	natFit:         15,
	pace:           8,
	stamina:        2,
	strength:       4,
	aggression:     12,
	anticipation:   11,
	bravery:        9,
	composure:      7,
	concentration:  11,
	decisions:      14,
	determination:  19,
	flair:          2,
	leadership:     12,
	offTheBall:     3,
	positioning:    12,
	teamwork:       7,
	vision:         6,
	workRate:       7,
	aerialReach:    12,
	commandOfArea:  12,
	communication:  8,
	eccentricity:   5,
	handling:       7,
	kicking:        7,
	oneOnOnes:      6,
	punching:       11,
	reflexes:       14,
	rushingOut:     6,
	throwing:       10,
}

var expectedHarry = Player{
	name:           "Harry Brumwell",
	age:            18,
	position:       "D (C)",
	corners:        5,
	crossing:       1,
	dribbling:      1,
	finishing:      2,
	firstTouch:     5,
	freeKickTaking: 1,
	heading:        6,
	longShots:      5,
	throwIns:       1,
	marking:        7,
	passing:        3,
	penaltyTaking:  1,
	tackling:       8,
	technique:      4,
	acceleration:   7,
	agility:        8,
	balance:        7,
	jumping:        7,
	natFit:         15,
	pace:           8,
	stamina:        6,
	strength:       5,
	aggression:     11,
	anticipation:   5,
	bravery:        7,
	composure:      6,
	concentration:  5,
	decisions:      4,
	determination:  12,
	flair:          5,
	leadership:     10,
	offTheBall:     5,
	positioning:    16,
	teamwork:       6,
	vision:         1,
	workRate:       7,
	aerialReach:    3,
	commandOfArea:  2,
	communication:  3,
	eccentricity:   4,
	handling:       3,
	kicking:        1,
	oneOnOnes:      1,
	punching:       2,
	reflexes:       2,
	rushingOut:     1,
	throwing:       2,
}
