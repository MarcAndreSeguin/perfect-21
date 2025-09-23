package main

import (
	"fmt"
	"math/rand"
)

type card struct {
	value int // 1..13
	suit  int // 0♥,1♦,2♣,3♠
}

func (card card) GetString() string {
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
	if num == 0  {
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
	g.dealerCards = make([]card, 1)
	g.playerCards[0] = g.shoe.deal(1)[0]
	g.dealerCards[0] = g.shoe.deal(1)[0]
	g.playerCards[1] = g.shoe.deal(1)[0]
}

func main() {
	game := game{}
	game.dealUpCards()
	
	fmt.Printf("Dealer upcard: %+v\n", game.dealerCards[0].GetString())
	fmt.Printf("Player cards: %+v %+v\n", game.playerCards[0].GetString(), game.playerCards[1].GetString())
}

