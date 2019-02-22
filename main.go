package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type bets struct {
	value    float64
	position int
}

func main() {
	var t []table
	reader := bufio.NewReader(os.Stdin)
	t, err := buildTable(reader)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var i int
		d := genShoe(6)
		d = shuffle(d)
		cut := len(d) - (len(t) * 5)
		for i < cut {
			var b []float64
			for p := range t {
				var tb float64
				t[p].player, tb, err = placeBet(reader, t[p].player)
				if err != nil {
					log.Fatal(err)
				}
				b = append(b, tb)
			}
			var dc []card
			i, dc = dealHands(t, d, i)
			for p := range t {
				var np player
				np, i = playerDecisions(reader, t[p].player, d, i, dc)
				t[p].player = np
			}
			dc, i = dealerLogic(dc, d, i)
			ds := sumHand(dc)
			sort.Ints(ds[:])
			var dv int
			if ds[0] != 0 {
				if ds[1] > 21 {
					dv = ds[0]
				}
				if ds[1] < 21 {
					dv = ds[1]
				}
			}
			if ds[0] == 0 {
				dv = ds[1]
			}
			for j := range t {
				po := payout(dv, t[j].player.cards, b[j])
				t[j].player.money = t[j].player.money + po
			}
			for p := range t {
				fmt.Println("\n", t[p].player.name, "has $", t[p].player.money, "available")
				t[p].player.cards = []card{}
			}
			for p := range t {
				if t[p].player.money == 0 {
					t[p].player, err = addFunds(reader, t[p].player)
					if err != nil {
						log.Fatal(err)
					}
					if t[p].player.money == 0 {
						fmt.Printf("Thanks for playing %s !", t[p].player.name)
						t = append(t[:p], t[p+1:]...)
					}
				}
			}
			if len(t) == 0 {
				os.Exit(0)
			}
		}
		fmt.Println("End of shoe, continue playing? [y,n]")
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		in = strings.TrimRight(in, "\r\n")
		if in == "n" {
			break
		}
	}
}
