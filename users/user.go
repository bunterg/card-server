package users

import (
	"errors"
	"time"
)

// User indicate type of card
type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

// ErrNotFound is used when a Room could not be found.
var ErrNotFound = errors.New("User not found")

// Repository provides access to the review storage.
type Repository interface {
	// GetAll returns a list of all reviews for a given Room ID.
	Get(int) (User, error)
	GetAll() []User
	// Add saves a given review.
	Add(User) error
}
