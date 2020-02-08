// test
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type card struct {
	num  int
	suit int
	val  int
}

type hand []card

type deck []card

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	startGame(deck)
	fmt.Println("End of program.")
}

func startGame(deck deck) {
	var dhand hand
	var phand hand
	var input string
	fmt.Println("Dealing starting hand:")
	phand = append(phand, deck.dealCard())
	dhand = append(dhand, deck.dealCard())
	phand = append(phand, deck.dealCard())
	dhand = append(dhand, deck.dealCard())
loop:
	for {
		fmt.Print("Your cards: ")
		phand.sayHand()
		fmt.Print("Worth ", phand.evalHand())
		fmt.Printf("\n")
		fmt.Println("Dealer showing:", dhand[1].nameCard())
		fmt.Print("Would you like to 'hit' or 'stay'?:")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		switch input {
		case "hit":
			phand = append(phand, deck.dealCard())
			fmt.Println("Adding a card to your hand")
		case "stay":
			fmt.Println("You have chosen to stay")
			break loop
		default:
			fmt.Println("I'm sorry I didn't understand that.")
		}
	}
}

//Take a card from the deck and return it.
func (deck *deck) dealCard() card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}

//Evaluate the hand's total value
func (h hand) evalHand() int {
	total := 0
	for i := range h {
		total += h[i].val
	}
	return total
}

//Shuffle the Deck
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

//Create the deck and return as a slice of cards.
func createDeck() deck {
	deck := []card{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			switch {
			case j == 0:
				deck = append(deck, card{num: j, suit: i, val: 11})
			case j > 9:
				deck = append(deck, card{num: j, suit: i, val: 10})
			default:
				deck = append(deck, card{num: j, suit: i, val: j + 1})
			}
		}
	}
	return deck
}

//Converts cards from int values to strings for output.
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

//Takes a hand and prints the contents to user.
func (hand hand) sayHand() {
	for i := range hand {
		fmt.Print(hand[i].nameCard())
		if i == len(hand)-1 {
			break
		}
		fmt.Printf(" and ")
	}
	fmt.Printf("\n")
}
