package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

//Hit adds a card to players hand
func Hit(p Player, d []Card, i int) (Player, int) {
	p.cards = append(p.cards, Card{value: d[i].value, suit: d[i].suit, face: d[i].face})
	i++
	return p, i
}

// SumHand will take the sum of a players current hand
func SumHand(h []Card) [2]int {
	var v [2]int
	var b bool
	for _, c := range h {
		if len(c.value) == 2 {
			v[0] += c.value[0]
			v[1] += c.value[1]
			b = true
		}
		if len(c.value) == 1 {
			v[0] += c.value[0]
			v[1] += c.value[0]
		}
	}
	if b == false {
		v[1] = 0
	}
	return v
}

//Table struct holds table information
type Table struct {
	position int
	player   Player
}

//BuildTable will add up to 6 players to start a game
func BuildTable(scan *bufio.Reader) ([]Table, error) {
	var t []Table
	fmt.Println("Add First Player")
	p, err := AddPlayer(scan)
	if err != nil {
		return t, err
	}
	t = append(t, Table{position: 1, player: p})
	i := 1
	for i < 6 {
		fmt.Println("Add Player? [y/n]")
		n, err := scan.ReadString('\n')
		if err != nil {
			return t, err
		}
		n = strings.ToLower(n)
		n = strings.TrimRight(n, "\r\n")
		if n == "n" {
			break
		}
		if n == "y" {
			i++
			p, err = AddPlayer(scan)
			t = append(t, Table{position: i, player: p})
		}
	}
	return t, nil
}

//DealHands will run initial hands to players
func DealHands(t []Table, d []Card, i int) (int, []Card) {
	var dc []Card
	for j := 0; j < 2; j++ {
		for k := range t {
			t[k].player.cards = append(t[k].player.cards, Card{value: d[i].value, suit: d[i].suit, face: d[i].face})
			i++
		}
		dc = append(dc, Card{value: d[i].value, suit: d[i].suit, face: d[i].face})
		i++
	}
	return i, dc
}

//Payout determines payout per player
func Payout(dv int, pc []Card, pb float64) float64 {
	var po float64
	pa := SumHand(pc)
	sort.Ints(pa[:])
	var pv int
	if pa[0] == 0 {
		pv = pa[1]
	}
	if pa[0] != 0 {
		if pa[1] < 21 {
			pv = pa[1]
		}
		if pa[1] > 21 {
			pv = pa[0]
		}
	}
	switch p := pv; {
	case p == 21 && len(pc) == 2:
		po = pb + pb*1.5
		fmt.Println("Blackjack!")
	case p > 21:
		po = 0
		fmt.Println("Bust!")
	case p == dv:
		po = pb
		fmt.Println("Draw")
	case p <= 21 && dv > 21:
		po = pb * 2
		fmt.Println("You win")
	case p > dv && dv <= 21:
		po = pb * 2
		fmt.Println("You win")
	case p < dv && dv <= 21:
		po = 0
		fmt.Printf("You lose - Dealer had %d, You had %d", dv, p)
	}
	return po
}
