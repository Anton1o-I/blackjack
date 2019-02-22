package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

//Hit adds a card to players hand
func hit(p player, d []card, i int) (player, int) {
	p.cards = append(p.cards, card{value: d[i].value, suit: d[i].suit, face: d[i].face})
	i++
	return p, i
}

// sumHand will take the sum of a players current hand
func sumHand(h []card) [2]int {
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
type table struct {
	position int
	player   player
}

//BuildTable will add up to 6 players to start a game
func buildTable(scan *bufio.Reader) ([]table, error) {
	var t []table
	fmt.Println("Add First Player")
	p, err := addPlayer(scan)
	if err != nil {
		return t, err
	}
	t = append(t, table{position: 1, player: p})
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
			p, err = addPlayer(scan)
			t = append(t, table{position: i, player: p})
		}
	}
	return t, nil
}

//dealHands will run initial hands to players
func dealHands(t []table, d []card, i int) (int, []card) {
	var dc []card
	for j := 0; j < 2; j++ {
		for k := range t {
			t[k].player.cards = append(t[k].player.cards, card{value: d[i].value, suit: d[i].suit, face: d[i].face})
			i++
		}
		dc = append(dc, card{value: d[i].value, suit: d[i].suit, face: d[i].face})
		i++
	}
	return i, dc
}

//payout determines payout per player
func payout(dv int, pc []card, pb float64) float64 {
	var po float64
	pa := sumHand(pc)
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
