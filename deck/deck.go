package deck

import (
	"fmt"
	"math/rand"
	"time"
)

// CardHandler stores methods that can act on decks.
type CardHandler interface {
	Shuffle()
	Draw()
}

// CardOptions stores rank and suit information for a specific game.
type CardOptions struct {
	ranks []string // rank is the name of the card ie. "two", "king", etc.
	suits []string
}

// Card stores information about specific cards.
type Card struct {
	rank  string
	suit  string
	value []int
}

// Deck contains a slice of card types.
type Deck struct {
	cards []Card
}

// BuildDeck will generate a deck that matches the
// required deck composition given by CardOptions.
func (c *CardOptions) BuildDeck() Deck {
	var d Deck
	for _, s := range c.suits {
		for _, r := range c.ranks {
			d.cards = append(d.cards, Card{rank: r, suit: s})
		}
	}
	return d
}

// Shuffle will randomly shuffle a deck.
func (d *Deck) Shuffle() {
	ld := len(d.cards)
	dest := make([]Card, ld)
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(ld)
	for i, v := range perm {
		dest[v] = d.cards[i]
	}
	fmt.Println(dest)
	d.cards = dest
}

// Draw returns the next card in the deck
func (d *Deck) Draw() Card {
	n, l := d.cards[0], d.cards[1:]
	d.cards = l
	return n
}
