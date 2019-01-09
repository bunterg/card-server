package cards

import (
	"errors"
	"time"
)

// CardType indicate type of card
type CardType int

// Card data model for game card profile
type Card struct {
	ID      int       `json:"id"`
	Att     int       `json:"att"`
	Def     int       `json:"def"`
	Cost    int       `json:"cost"`
	Created time.Time `json:"created"`
	Class   CardType  `json:"class"`
}

const (
	basic CardType = iota
	special
	legendary
)

func (cardType CardType) String() string {
	names := [...]string{
		"Basic",
		"Special",
		"Legendary"}

	// invalid cardType, out of range
	if cardType < basic || cardType > legendary {
		return "Unknown"
	}

	return names[cardType]
}

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("Card not found")
var ErrDuplicate = errors.New("Card Already Exists")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given beer ID.
	Get(int) (Card, error)
	GetAll() []Card
	// Add saves a given review.
	Add(Card) error
}
