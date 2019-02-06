package cards

import (
	"errors"
	"time"
)

// CardType indicate type of card
type CardType int

// Color indicate color of card
type Color int

// Card data model for game card profile
type Card struct {
	ID      int       `json:"id"`
	Number  int       `json:"number"`
	Created time.Time `json:"created"`
	Class   CardType  `json:"class"`
	Color   Color     `json:"color"`
}

const (
	Number CardType = iota
	Skip
	Reverse
	DrawTwo
	DrawFour
	Wild
)

const (
	Blue Color = iota
	Green
	Red
	Yellow
	Uncolor
)

func (c CardType) String() string {
	names := [...]string{
		"Number",
		"Skip",
		"Reverse",
		"DrawTwo",
		"DrawFour",
		"Wild"}

	// invalid cardType, out of range
	if c < Number || c > Wild {
		return "Unknown"
	}

	return names[c]
}

func (c Color) String() string {
	names := [...]string{
		"Blue",
		"Green",
		"Red",
		"Yellow",
		"Uncolor"}

	// invalid cardType, out of range
	if c < Blue || c > Uncolor {
		return "Unknown"
	}

	return names[c]
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
