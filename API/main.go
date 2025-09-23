package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"math"
	"embed"
)

//go:embed perfect21-strategy.json
var f embed.FS


type card struct {
	value int // 1..13
	suit  int // 0♥,1♦,2♣,3♠
}

func (card card) getString() string {
	var suit, value string
	switch card.suit {
	case 0:
		suit = "♥"
	case 1:
		suit = "♦"
	case 2:
		suit = "♣"
	case 3:
		suit = "♠"
	}
	switch card.value {
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	case 1:
		value = "A"
	default:
		value = fmt.Sprintf("%d", card.value)
	}
	return value + suit
}

func getCardValue(card card) int {
	return int(math.Min(10, float64(card.value)))
}

type deck struct{ cards []card }

func (d *deck) create() {
	newDeck := make([]card, 0, 52)
	for s := 0; s < 4; s++ {
		for v := 1; v <= 13; v++ {
			newDeck = append(newDeck, card{value: v, suit: s})
		}
	}
	d.cards = newDeck
}

type shoe struct {
	cards []card
}

const decksPerShoe = 6

func (s *shoe) create() {
	newShoe := make([]card, 0, decksPerShoe*52)
	for i := 0; i < decksPerShoe; i++ {
		var d deck
		d.create()
		newShoe = append(newShoe, d.cards...)
	}
	rand.Shuffle(len(newShoe), func(i, j int) { newShoe[i], newShoe[j] = newShoe[j], newShoe[i] })
	s.cards = newShoe

	// burn one card at the start of a fresh shoe
	s.deal(1)
}

func (s *shoe) deal(num uint) []card {
	if num == 0 {
		return nil
	}
	n := int(num)

	dealt := s.cards[:n]
	s.cards = s.cards[n:]

	return dealt
}

type game struct {
	shoe        shoe
	playerCards []card
	dealerCards []card
}

func (g *game) dealUpCards() {
	g.shoe.create()

	g.playerCards = make([]card, 2)
	g.dealerCards = make([]card, 2)

	// Player
	g.playerCards[0] = g.shoe.deal(1)[0]
	// Dealer
	g.dealerCards[0] = g.shoe.deal(1)[0]
	// Player
	g.playerCards[1] = g.shoe.deal(1)[0]
	// Dealer
	g.dealerCards[1] = g.shoe.deal(1)[0]

}

// returns: hand total, isSoft, isBlackJack
func evaluateHand(h []card) (int, bool, bool, bool) {
	aces := 0
	total := 0
	isSoft := false

	for _, card := range h {
		if card.value == 1 {
			aces++
			continue
		}
		v := card.value
		if v > 10 {
			v = 10
		}
		total += v
	}

	// count all aces as 1, then promote one to 11 if it fits
	total += aces
	if aces > 0 && total+10 <= 21 {
		total += 10
		isSoft = true
	}

	isBlackjack := total == 21 && len(h) == 2
	isTenVal := func(v int) bool { return v == 10 || v == 11 || v == 12 || v == 13 }
	isPair := len(h) == 2 && (h[0].value == h[1].value || (isTenVal(h[0].value) && isTenVal(h[1].value)))
	return total, isSoft, isBlackjack, isPair
}

func isBlackJack(h []card) bool {
	_, _, bj, _ := evaluateHand(h)
	return bj
}

func getCardsValue(h []card) int {
	total, _, _, _ := evaluateHand(h)
	return total
}

func isSoft(h []card) bool {
	_, soft, _, _ := evaluateHand(h)
	return soft
}

func isPair(h []card) bool {
	_, _, _, pair := evaluateHand(h)
	return pair
}

// BuildScenario constructs the JSON-ready payload from the current game state.
// Assumes you already have: isBlackJack, isSoft, getCardsValue, and card.getString().
func BuildScenario(g *game) Scenario {
	return Scenario{
		Dealer: DealerInfo{
			PrettyString: []string{g.dealerCards[0].getString(), "?"}, // show "?" in trainer mode
			UpCardValue:       getCardValue(g.dealerCards[0]),
			HoleCardValue:     getCardValue(g.dealerCards[1]),
			IsBlackJack:  isBlackJack(g.dealerCards),
		},
		Player: PlayerInfo{
			PrettyString: []string{
				g.playerCards[0].getString(),
				g.playerCards[1].getString(),
			},
			Card1Value:       getCardValue(g.playerCards[0]),
			Card2Value:       getCardValue(g.playerCards[1]),
			PlayerTotal: getCardsValue(g.playerCards),
			IsPair:      isPair(g.playerCards),
			IsSoft:      isSoft(g.playerCards),
			IsBlackJack: isBlackJack(g.playerCards),
		},
		CorrectAction: determineCorrectAction(g), // TODO: replace with determineCorrectAction(...) later
	}
}

func determineCorrectAction(g *game) CorrectAction {
	// TO DO: Lookup the perfect strategy in JSON file

	// dealer BJ or player BJ: NONE
	if isBlackJack(g.dealerCards) || isBlackJack(g.playerCards) {
		return ActionNone
	}

	// player pair
	if isPair(g.playerCards) { // TO DO
	} else if isSoft(g.playerCards) { // TO DO
	} else { // TO DO
	}
	return ActionNone // placeholder return to avoid errors
}

func main() {
	g := game{}
	g.dealUpCards()

	sc := BuildScenario(&g)
	b, err := json.MarshalIndent(sc, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	data, _ := f.ReadFile("perfect21-strategy.json")
	print(string(data))
}

//-- JSON objects required //

// Top-level payload
type Scenario struct {
	Dealer        DealerInfo    `json:"dealer"`
	Player        PlayerInfo    `json:"player"`
	CorrectAction CorrectAction `json:"correctAction"`
}

// Dealer sub-object
type DealerInfo struct {
	PrettyString []string `json:"prettyString"` // e.g. ["8♦","?"]
	UpCardValue       int      `json:"upCardValue"`       // 2..11 (Ace = 11 or 1; choose your convention)
	HoleCardValue     int      `json:"holeCardValue"`     // hidden value if you precompute; else 0
	IsBlackJack  bool     `json:"isBlackJack"`
}

// Player sub-object
type PlayerInfo struct {
	PrettyString []string `json:"prettyString"` // e.g. ["10♦","4♣"]
	Card1Value        int      `json:"card1Value"`
	Card2Value        int      `json:"card2Value"`
	PlayerTotal  int      `json:"playerTotal"` // computed hand total (respecting soft rules)
	IsPair       bool     `json:"isPair"`
	IsSoft       bool     `json:"isSoft"`
	IsBlackJack  bool     `json:"isBlackJack"`
}

type CorrectAction string

const (
	ActionHit       CorrectAction = "HIT"
	ActionStand     CorrectAction = "STAND"
	ActionDouble    CorrectAction = "DOUBLE"
	ActionSplit     CorrectAction = "SPLIT"
	ActionSurrender CorrectAction = "SURRENDER"
	ActionNone      CorrectAction = "NONE"
)
