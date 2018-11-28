package decks

import (
	"card-server/cards"
	"errors"
	"time"
)

// Deck indicate type of card
type Deck struct {
	ID      int
	Cards   []cards.Card
	Created time.Time
}

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("Deck not found")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given beer ID.
	Get(int) []Deck
	GetAll() []Deck
	// Add saves a given review.
	Add(Deck) error
}
