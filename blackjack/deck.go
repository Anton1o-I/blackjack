package main

import (
	"math/rand"

	"github.com/pborman/uuid"
)

// card is the struct for a single card in a deck
type card struct {
	id    string
	value []int
	face  string
	suit  string
}

var cardValues = map[string][]int{
	"ace":   []int{11, 1},
	"king":  []int{10},
	"queen": []int{10},
	"jack":  []int{10},
	"10":    []int{10},
	"9":     []int{9},
	"8":     []int{8},
	"7":     []int{7},
	"6":     []int{6},
	"5":     []int{5},
	"4":     []int{4},
	"3":     []int{3},
	"2":     []int{2},
}

var suits = []string{"clubs", "spades", "hearts", "diamonds"}

func genDeck() []card {
	var deck []card
	for _, s := range suits {
		for f, val := range cardValues {
			c := card{id: uuid.NewUUID().String(), value: val, face: f, suit: s}
			deck = append(deck, c)
		}
	}
	return deck
}

//genShoe builds a shoe with num amount of decks
func genShoe(num int) []card {

	var shoe []card
	for i := 0; i < num; i++ {
		shoe = append(shoe, genDeck()...)
	}
	return shoe
}

// shuffle will shuffle a slice of Cards
func shuffle(s []card) []card {
	dest := make([]card, len(s))
	perm := rand.Perm(len(s))
	for i, v := range perm {
		dest[v] = s[i]
	}
	return dest
}
