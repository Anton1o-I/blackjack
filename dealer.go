package main

import (
	"fmt"
	"sort"
)

//DealerLogic determines the logic the dealer uses to stand or get another card
func DealerLogic(h []Card, d []Card, i int) ([]Card, int) {
	v := SumHand(h)
	sort.Ints(v[:])
	var s bool
	for {
		fmt.Println(v)
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
			h = append(h, d[i])
			v = SumHand(h)
			sort.Ints(v[:])
			i++
		}
	}
	return h, i
}
