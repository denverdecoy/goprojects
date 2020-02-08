// test
package main

import (
	"fmt"
	"math/rand"
	//"strings"
	"time"
)

type card struct {
	num  int
	suit int
	val  int
}

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	startGame(deck)
}

func startGame(deck []card) {
	var dhand []card
	var phand []card
	//var input string
	fmt.Println("Dealing starting hand:")
	for i := 0; i < 2; i++ {
		phand = append(phand, deck[len(deck)-1])
		deck = deck[:len(deck)-1]
		dhand = append(dhand, deck[len(deck)-1])
		deck = deck[:len(deck)-1]
	}
	fmt.Println("Your hand:", phand[0].nameCard(), "and", phand[1].nameCard())
	fmt.Println("Dealer is showing:", dhand[1].nameCard())
}

func shuffleDeck(d []card) []card {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}

func createDeck() []card {
	deck := []card{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			deck = append(deck, card{num: j, suit: i, val: j + 1})
		}
	}
	return deck
}

func (c card) nameCard() string {
	s1 := ""
	s2 := ""
	if c.num < 1 {
		s1 = fmt.Sprint("Ace")
	} else if c.num < 10 {
		s1 = fmt.Sprint(c.num + 1)
	} else {
		switch c.num {
		case 10:
			s1 = fmt.Sprint("Jack")
		case 11:
			s1 = fmt.Sprint("Queen")
		case 12:
			s1 = fmt.Sprint("King")
		}
	}
	switch c.suit {
	case 0:
		s2 = " of Hearts"
	case 1:
		s2 = " of Diamonds"
	case 2:
		s2 = " of Clubs"
	case 3:
		s2 = " of Spades"
	}
	final := fmt.Sprint(s1, s2)
	return final
}
