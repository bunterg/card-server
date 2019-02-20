package decks

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bunterg/card-server/cards"
)

// Deck indicate type of card
type Deck struct {
	ID      int          `json:"id"`
	Cards   []cards.Card `json:"cards"`
	Created time.Time    `json:"created"`
}

// Draw first card from the deck
func (d Deck) Draw() cards.Card {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}

// Play a card into deck
func (d Deck) Play(c cards.Card) {
	d.Cards = append(d.Cards, c)
}

// LastCard return current card
func (d Deck) LastCard() cards.Card {
	return d.Cards[len(d.Cards)-1]
}

// Combine cards from current deck with another deck cards
func (d Deck) Combine(d2 Deck) {
	d.Cards = append(d.Cards, d2.Cards...)
}

// Shuffle cards from the deck
func (d Deck) Shuffle() {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// ErrNotFound is used when a deck could not be found.
var ErrNotFound = errors.New("Deck not found")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given deck ID.
	Get(int) (Deck, error)
	GetAll() []Deck
	// Add saves a given review.
	Add(Deck) error
}
