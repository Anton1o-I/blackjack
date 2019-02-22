package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//Player struct holds information on a single player
type Player struct {
	name  string
	money float64
	cards []Card
}

//AddPlayer function to add a player to a game
func AddPlayer(scan *bufio.Reader) (Player, error) {
	p := Player{}
	fmt.Println("Player Name:")
	n, err := scan.ReadString('\n')
	if err != nil {
		return p, err
	}
	n = strings.TrimRight(n, "\r\n")
	fmt.Println("Buy in:")
	b, err := scan.ReadString('\n')
	if err != nil {
		return p, err
	}
	b = strings.TrimRight(b, "\r\n")
	bf, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return p, err
	}
	p = Player{name: n, money: bf}
	return p, nil
}

func wager(scan *bufio.Reader, p Player) (float64, error) {
	var w float64
	fmt.Printf("%s, how much would you like to wager?\n", p.name)
	v, err := scan.ReadString('\n')
	if err != nil {
		return w, err
	}
	v = strings.TrimRight(v, "\r\n")
	vf, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return w, err
	}
	w = vf
	return w, err
}

//PlaceBet takes a user input bet
func PlaceBet(scan *bufio.Reader, p Player) (Player, float64, error) {
	var b float64
	for {
		wf, err := wager(scan, p)
		if err != nil {
			return p, b, err
		}
		for p.money < wf {
			fmt.Println("You do not have that much money")
			w, err := wager(scan, p)
			if err != nil {
				return p, b, err
			}
			wf = w
		}
		b = wf
		p.money -= b
		break
	}
	return p, b, nil
}

//PlayerDecisions allows player to decide to hit or stand and makes calculations
func PlayerDecisions(scan *bufio.Reader, p Player, d []Card, i int, dc []Card) (Player, int) {
	for {
		o, err := HandOptions(scan, p, dc)
		if err != nil {
			log.Fatal(err)
		}
		if o == "stand" {
			break
		}
		if o == "hit" {
			p, i = Hit(p, d, i)
		}
		cv := SumHand(p.cards)
		sort.Ints(cv[:])
		var bust bool
		if cv[0] != 0 {
			if cv[0] > 21 && cv[1] > 21 {
				bust = true
			}
		}
		if cv[0] == 0 {
			if cv[1] > 21 {
				bust = true
			}
		}
		if bust {
			fmt.Printf("Bust %v ", SumHand(p.cards))
			break
		}
	}
	fmt.Println(p.name, "has", p.cards, "for total value of", SumHand(p.cards))
	return p, i
}

//HandOptions are the user input for decisions they can make in the hand
func HandOptions(scan *bufio.Reader, p Player, dc []Card) (string, error) {
	options := []string{"stand", "hit"}
	var valid bool
	var n string
	for valid == false {
		var err error
		fmt.Printf("%s what do you want to do?\n Current Cards %v \nCurrent Value %v \n [Stand, Hit]\n Dealer Showing Card %v (value %v) \n",
			p.name,
			p.cards,
			SumHand(p.cards),
			dc[0].face,
			dc[0].value,
		)
		n, err = scan.ReadString('\n')
		if err != nil {
			return "", err
		}
		n = strings.ToLower(n)
		n = strings.TrimRight(n, "\r\n")
		for _, o := range options {
			if n == o {
				valid = true
			}
		}
		if valid == true {
			break
		}
		fmt.Printf("%s not a recognized option, try again \n\n", n)
	}
	return n, nil
}

//AddFunds checks if a player wants to add money
func AddFunds(scan *bufio.Reader, p Player) (Player, error) {
	fmt.Println("You're out of money, do you want to add funds? [y,n]")
	i, err := scan.ReadString('\n')
	if err != nil {
		return p, err
	}
	i = strings.TrimRight(i, "\r\n")
	if i == "n" {
		return p, nil
	}
	if i == "y" {
		fmt.Println("How much do you want to add?")
		m, err := scan.ReadString('\n')
		if err != nil {
			return p, err
		}
		m = strings.TrimRight(i, "\r\n")
		mf, err := strconv.ParseFloat(m, 64)
		if err != nil {
			return p, err
		}
		p.money += mf
	}
	return p, nil
}
