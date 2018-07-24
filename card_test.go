package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Four, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Four of Diamonds
	// Jack of Spades
	// Nine of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	for _, c := range cards {
		if c.Suit == Joker {
			t.Error("Standard deck shouldn't include any Jokers")
		}
	}

	if len(cards) != 13*4 {
		fmt.Println("found", len(cards), "cards in the standard deck")
		t.Error("Wrong number of cards in the deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected 'Ace of Spades' to be the first card, Recieved:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	defaultCards := New(DefaultSort)
	shuffledCards := Shuffle(defaultCards)
	counter := 0
	for i := 0; i < len(defaultCards); i++ {
		if defaultCards[i] == shuffledCards[i] {
			counter++
		}
	}
	if counter > 50 {
		t.Error("Cards don't appear to have been shuffled")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(5))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 5 {
		t.Error("Expected 5 Jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all 2s and 3s to be filtered out, received:", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 52*3 {
		t.Error("Expected 3 decks (52*3 cards), received:", len(cards))
	}
}
