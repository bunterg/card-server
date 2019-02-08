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
	Owner   users.User   `json:"Owner"`
	Created time.Time    `json:"created"`
}

// ErrNotFound is used when a Room could not be found.
var ErrNotFound = errors.New("Room not found")

// ErrRoomFull is used when a Room is full
var ErrRoomFull = errors.New("Room aleardy full")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given beer ID.
	Get(int) (Room, error)
	GetAll() []Room
	// Add saves a given review.
	Add(Room) error
	AddPlayer(Room, users.User) error
}
