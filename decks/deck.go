package decks

import (
	"errors"
	"time"

	"github.com/bunterg/card-server/cards"
)

// Deck indicate type of card
type Deck struct {
	ID      int          `json:"id"`
	Cards   []cards.Card `json:"cards"`
	Created time.Time    `json:"created"`
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
