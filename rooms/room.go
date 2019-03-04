package rooms

import (
	"errors"
	"time"

	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/decks"
	"github.com/bunterg/card-server/users"
)

// Room player game lobby
type Room struct {
	ID      int          `json:"id"`
	Owner   users.User   `json:"Owner"`
	Created time.Time    `json:"created"`
	Players []users.User `json:"players"`
	Board   Board        `json:"board"`
}

// Board playars and cards
type Board struct {
	Graveyard decks.Deck `json:"gravetard"`
	Deck      decks.Deck `json:"deck"`
	Turn      int        `json:"turn"`
	Hands     []Hand     `json:"hands"`
}

// Hand player cards
type Hand struct {
	Cards  []cards.Card `json:"cards"`
	Player users.User   `json:"player"`
}

// ErrNotFound is used when a Room could not be found.
var ErrNotFound = errors.New("Room not found")

// ErrRoomFull is used when a Room is full
var ErrRoomFull = errors.New("Room aleardy full")

// ErrNotEnoughPlayers is used when a Room is full
var ErrNotEnoughPlayers = errors.New("Room needs more players")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given Room ID.
	Get(int) (Room, error)
	GetAll() []Room
	// Add saves a given review.
	Add(Room, users.User) (Room, error)
	AddPlayer(Room, users.User) error
}
