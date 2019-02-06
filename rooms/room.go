package rooms

import (
	"errors"
	"time"

	"github.com/bunterg/card-server/decks"
	"github.com/bunterg/card-server/users"
)

// Room indicate type of card
type Room struct {
	ID      int          `json:"id"`
	Deck    decks.Deck   `json:"deck"`
	Users   []users.User `json:"users"`
	Created time.Time    `json:"created"`
}

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("Room not found")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given beer ID.
	Get(int) (Room, error)
	GetAll() []Room
	// Add saves a given review.
	Add(Room) error
}
