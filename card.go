//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit attribute of the Card (Ex. Heart)
type Suit uint8

// Valid Suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank attribute of the Card (Ex. Queen)
type Rank uint8

// Valid Ranks
const (
	_ Rank = iota
	Ace
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

const (
	minRank = Ace
	maxRank = King
)

// Card represents a single card in the deck
type Card struct {
	Suit
	Rank
}

// String returns an appropriate string to represent the Card
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New generates the card deck
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort takes in a deck and returns it after sorting by Suit and Rank
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort takes in a deck and returns it after sorting
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less determines if the value of card 'i' is less than the value of card 'j'
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Shuffle will mix up the order of the cards in the deck
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers adds a number 'n' of Card to a deck
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

// Filter takes in a filter function and returns a filtered deck
func Filter(f func(card Card) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		ret := make([]Card, len(cards))
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck returns a function that duplicates the passed in deck 'n' number of times, returning 'n' number of decks as one deck
func Deck(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
