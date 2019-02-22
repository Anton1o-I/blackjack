package main

import (
	"fmt"
	"sort"
)

//dealerLogic determines the logic the dealer uses to stand or get another card
func dealerLogic(h []card, d []card, i int) ([]card, int) {
	v := sumHand(h)
	sort.Ints(v[:])
	var s bool
	for {
		fmt.Println("Dealer has", h, v)
		if v[0] != 0 {
			if v[0] > 16 && v[1] > 16 {
				s = true
			}
		}
		if v[0] == 0 {
			if v[1] > 16 {
				s = true
			}
		}
		if s == true {
			break
		}
		if s != true {
			h = append(h, card{face: d[i].face, value: d[i].value})
			v = sumHand(h)
			sort.Ints(v[:])
			i++
		}
	}
	return h, i
}
