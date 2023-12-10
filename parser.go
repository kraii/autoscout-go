package main

import (
	"autoscout-go/queue"
	"bufio"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html"
	"os"
	"strconv"
)

type Player struct {
	name     string
	age      int
	position string

	// technical
	corners        int
	crossing       int
	dribbling      int
	finishing      int
	firstTouch     int
	freeKickTaking int
	heading        int
	longShots      int
	throwIns       int
	marking        int
	passing        int
	penaltyTaking  int
	tackling       int
	technique      int

	// physical
	acceleration int
	agility      int
	balance      int
	jumping      int
	natFit       int
	pace         int
	stamina      int
	strength     int

	// mental
	aggression    int
	anticipation  int
	bravery       int
	composure     int
	concentration int
	decisions     int
	determination int
	flair         int
	leadership    int
	offTheBall    int
	positioning   int
	teamwork      int
	vision        int
	workRate      int

	// Goalkeeping
	aerialReach   int
	commandOfArea int
	communication int
	eccentricity  int
	handling      int
	kicking       int
	oneOnOnes     int
	punching      int
	reflexes      int
	rushingOut    int
	throwing      int
}

func Parse(filename string) ([]*Player, error) {

	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	rows := findElements(doc, "tr")

	headerIndexes := make(map[string]int)
	for i, header := range findElements(doc, "th") {
		cellText := getText(header)
		_, existing := headerIndexes[cellText]
		if existing {
			// Attributes have duplicate heading names in the html table
			headerIndexes[spew.Sprintf("%s2", cellText)] = i
		} else {
			headerIndexes[cellText] = i
		}
	}

	result := make([]*Player, len(rows)-1)
	for i, row := range rows[1:] {
		columns := findElements(row, "td")
		p := Player{
			name:     getAttr("Name", headerIndexes, columns),
			age:      getAttrI("Age", headerIndexes, columns),
			position: getAttr("Position", headerIndexes, columns),

			corners:        getAttrI("Cor", headerIndexes, columns),
			crossing:       getAttrI("Cro", headerIndexes, columns),
			dribbling:      getAttrI("Dri", headerIndexes, columns),
			finishing:      getAttrI("Fin", headerIndexes, columns),
			firstTouch:     getAttrI("Fir", headerIndexes, columns),
			freeKickTaking: getAttrI("Fre", headerIndexes, columns),
			heading:        getAttrI("Hea", headerIndexes, columns),
			longShots:      getAttrI("Lon", headerIndexes, columns),
			marking:        getAttrI("Mar", headerIndexes, columns),
			passing:        getAttrI("Pas", headerIndexes, columns),
			penaltyTaking:  getAttrI("Pen", headerIndexes, columns),
			tackling:       getAttrI("Tck", headerIndexes, columns),
			technique:      getAttrI("Tec", headerIndexes, columns),
			throwIns:       getAttrI("L Th", headerIndexes, columns),
			aggression:     getAttrI("Agg", headerIndexes, columns),
			anticipation:   getAttrI("Ant", headerIndexes, columns),
			bravery:        getAttrI("Bra", headerIndexes, columns),
			composure:      getAttrI("Cmp", headerIndexes, columns),
			concentration:  getAttrI("Cnt", headerIndexes, columns),
			decisions:      getAttrI("Dec", headerIndexes, columns),
			flair:          getAttrI("Fla", headerIndexes, columns),
			determination:  getAttrI("Det", headerIndexes, columns),
			leadership:     getAttrI("Ldr", headerIndexes, columns),
			offTheBall:     getAttrI("OtB", headerIndexes, columns),
			positioning:    getAttrI("Pos", headerIndexes, columns),
			teamwork:       getAttrI("Tea", headerIndexes, columns),
			vision:         getAttrI("Vis", headerIndexes, columns),
			workRate:       getAttrI("Wor", headerIndexes, columns),
			acceleration:   getAttrI("Acc", headerIndexes, columns),
			agility:        getAttrI("Agi", headerIndexes, columns),
			balance:        getAttrI("Bal", headerIndexes, columns),
			jumping:        getAttrI("Jum", headerIndexes, columns),
			natFit:         getNatFit(headerIndexes, columns),
			pace:           getAttrI("Pac", headerIndexes, columns),
			stamina:        getAttrI("Sta", headerIndexes, columns),
			strength:       getAttrI("Str", headerIndexes, columns),
			aerialReach:    getAttrI("Aer", headerIndexes, columns),
			commandOfArea:  getAttrI("Cmd", headerIndexes, columns),
			communication:  getAttrI("Com", headerIndexes, columns),
			eccentricity:   getAttrI("Ecc", headerIndexes, columns),
			handling:       getAttrI("Han", headerIndexes, columns),
			kicking:        getAttrI("Kic", headerIndexes, columns),
			oneOnOnes:      getAttrI("1v1", headerIndexes, columns),
			punching:       getAttrI("Pun", headerIndexes, columns),
			reflexes:       getAttrI("Ref", headerIndexes, columns),
			rushingOut:     getAttrI("TRO", headerIndexes, columns),
			throwing:       getAttrI("Thr", headerIndexes, columns),
		}
		result[i] = &p
	}
	return result, nil
}

func getNatFit(headerIndexes map[string]int, columns []*html.Node) int {
	result := getAttrI("Nat", headerIndexes, columns)
	// Workaround for nationality also being called Nat
	if result == -1 {
		return getAttrI("Nat1", headerIndexes, columns)
	}
	return result
}

func findElements(node *html.Node, tag string) []*html.Node {
	visitStack := queue.EmptyQueue[*html.Node]()
	visitStack.Push(node)
	found := make([]*html.Node, 0, 10)

	for !visitStack.IsEmpty() {
		current, _ := visitStack.Pop()

		if current.Type == html.ElementNode && current.Data == tag {
			found = append(found, current)
		}

		for child := current.FirstChild; child != nil; child = child.NextSibling {
			visitStack.Push(child)
		}
	}

	return found
}

func getAttrI(attr string, headerIndexes map[string]int, row []*html.Node) int {
	i, e := strconv.Atoi(getAttr(attr, headerIndexes, row))
	if e == nil {
		return i
	} else {
		return -1
	}
}

func getAttr(attr string, headerIndexes map[string]int, row []*html.Node) string {
	attrI, present := headerIndexes[attr]
	if !present {
		return "unknown"
	}
	return getText(row[attrI])
}

func getText(node *html.Node) string {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			return child.Data
		}
	}
	return ""
}
