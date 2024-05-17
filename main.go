package card

import (
	"fmt"
	"math/rand"
	"sort"
)

type CardType int
type CardValue int

const (
	Spades CardType = iota
	Diamonds
	Clubs
	Hearts
	Jokers
)

const (
	Ace CardValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	suit  CardType
	value CardValue
}

type Options struct {
	sort           func(i, j int) bool
	shuffle        bool
	numberOfJokers int
	omit           map[int]bool
	numberOfDecks  int
}

type Option func(*Options)

func WithSort(sortFunc func(i, j int) bool) Option {
	return func(o *Options) {
		o.sort = sortFunc
	}
}

func WithShuffle(isShuffle bool) Option {
	return func(o *Options) {
		o.shuffle = isShuffle
	}
}

func NumbersOfJokers(jokers int) Option {
	return func(o *Options) {
		o.numberOfJokers = jokers
	}
}

func WhatToOmit(cardsToEmit []int) Option {
	var omittedValues = make(map[int]bool)
	for _,v:=range cardsToEmit{
		omittedValues[v]=true
	}
	return func(o *Options) {
		o.omit = omittedValues
	}
}

func NumberOfDecks(decks int) Option {
	return func(o *Options) {
		o.numberOfDecks = decks
	}
}


func New(option ...Option) [][]Card {
	var decks [][]Card
	var deck []Card
	var suitOrder = []CardType{Spades, Diamonds, Clubs, Hearts}
	var valueOrder = []CardValue{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	options := Options{
		shuffle:        false,
		numberOfJokers: 0,
		numberOfDecks:  1,
	}

	for _, v := range option {
		v(&options)
	}

	if options.numberOfDecks<=0{
		return nil
	}

	// Adding number of jokers to the deck
	for range make([]int, options.numberOfJokers) {
		suitOrder = append(suitOrder, Jokers)
	}

	// Sorting suits depending upon passed option
	if options.sort != nil {
		sort.Slice(suitOrder, options.sort)
	}

	// Omitting specific cards
	if len(options.omit)>0{
		var tempValueOrder []CardValue
		for _,v:=range valueOrder{
			if options.omit[int(v)]{
				continue
			}
			tempValueOrder = append(tempValueOrder, v)
		}
		valueOrder=tempValueOrder
	}

	// Creating deck
	for _, suit := range suitOrder {
		for _, value := range valueOrder {
			if suit == Jokers {
				deck = append(deck, Card{
					suit: Jokers,
				})
				break
			}
			var card Card = Card{
				suit:  suit,
				value: value,
			}
			deck = append(deck, card)
		}
	}

	// Shuffling deck depending upon the option
	if options.shuffle {
		shuffleDeck(&deck)
	}

	// Generating more decks if asked
	if options.numberOfDecks>1{
		for range make([]int,options.numberOfDecks){
			tempDeck:=New(append(option,NumberOfDecks(options.numberOfDecks-1))...)
			if tempDeck!=nil{
				decks = append(decks, tempDeck[0])
			}
		}

		return decks
	}
	return append(decks,deck)
}

func shuffleDeck(deck *[]Card){
	var newDeck =[]Card{}
	var tempDeck []Card=*deck
	for range make([]int,len(*deck)){
		index:=rand.Intn(len((tempDeck)))
		newDeck = append(newDeck, tempDeck[index])
		tempDeck = append(tempDeck[:index],tempDeck[index+1:]...)
	}
	*deck = newDeck
}

func main() {
	fmt.Println(New(NumberOfDecks(1),NumbersOfJokers(3),WhatToOmit([]int{1,2,3,4,5,6,7,8,9,10})))
	test(struct{sort bool; shuffle bool; numberJokers int}{sort: true,shuffle: true, numberJokers: 2})
}

func test(a struct{
	sort bool
	shuffle bool
	numberJokers int
}){
	println(a.sort,a.numberJokers)
}

