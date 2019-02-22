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
	var t []Table
	reader := bufio.NewReader(os.Stdin)
	t, err := BuildTable(reader)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var i int
		d := GenShoe(6)
		d = Shuffle(d)
		cut := len(d) - (len(t) * 5)
		for i < cut {
			var b []float64
			for p := range t {
				var tb float64
				t[p].player, tb, err = PlaceBet(reader, t[p].player)
				if err != nil {
					log.Fatal(err)
				}
				b = append(b, tb)
			}
			var dc []Card
			i, dc = DealHands(t, d, i)
			for p := range t {
				var np Player
				np, i = PlayerDecisions(reader, t[p].player, d, i, dc)
				t[p].player = np
			}
			dc, i = DealerLogic(dc, d, i)
			ds := SumHand(dc)
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
				po := Payout(dv, t[j].player.cards, b[j])
				t[j].player.money = t[j].player.money + po
			}
			fmt.Println(t)
			for p := range t {
				t[p].player.cards = []Card{}
			}
			for p := range t {
				if t[p].player.money == 0 {
					t[p].player, err = AddFunds(reader, t[p].player)
					if err != nil {
						log.Fatal()
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
