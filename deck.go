package main

import (
	"math/rand"

	"github.com/pborman/uuid"
)

// Card is the struct for a single card in a deck
type Card struct {
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

func genDeck() []Card {
	var deck []Card
	for _, s := range suits {
		for f, val := range cardValues {
			c := Card{id: uuid.NewUUID().String(), value: val, face: f, suit: s}
			deck = append(deck, c)
		}
	}
	return deck
}

//GenShoe builds a shoe with num amount of decks
func GenShoe(num int) []Card {

	var shoe []Card
	for i := 0; i < num; i++ {
		shoe = append(shoe, genDeck()...)
	}
	return shoe
}

// Shuffle will shuffle a slice of Cards
func Shuffle(s []Card) []Card {
	dest := make([]Card, len(s))
	perm := rand.Perm(len(s))
	for i, v := range perm {
		dest[v] = s[i]
	}
	return dest
}
