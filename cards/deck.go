package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Card 牌
type Card struct {
	suit  string
	value string
}

func (c Card) print() {
	fmt.Print(c.value, " of ", c.suit)
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.value, c.suit)
}

// NewCard create to card
func NewCard(suit string, value string) Card {
	return Card{suit, value}
}

// Deck a whole set of cards
type Deck []Card

// NewDeck create a deck of cards
func newDeck() Deck {
	deck := Deck{}
	cardSuits := []string{"Spade", "Heart", "Club", "Diamond"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			c := Card{suit, value}
			deck = append(deck, c)
		}
	}
	return deck
}

func newDeckFromFile(filename string) Deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	deck := Deck{}
	deckStr := string(bytes)
	cardStrs := strings.Split(deckStr, ",")
	for _, cardStr := range cardStrs {
		cardSlice := strings.Split(cardStr, " of ")
		deck = append(deck, Card{cardSlice[1], cardSlice[0]})
	}
	return deck
}

func (d Deck) print() {
	for i, card := range d {
		// fmt.Printf("%d %s\n", i, card)
		fmt.Println(i, card)
	}
}

func (d Deck) toString() string {
	result := make([]string, len(d))
	for i, v := range d {
		result[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(result, ",")
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		// j := rand.Intn(i + 1)
		j := r.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}
