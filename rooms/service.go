package rooms

import (
	"github.com/bunterg/card-server/users"
)

// Service provides player adding operations
type Service interface {
	AddPlayer(Room, ...users.User)
}

type service struct {
	rR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(rR Repository) Service {
	return &service{rR}
}

// AddCard adds the given user(s) to the database
func (s *service) AddPlayer(r Room, b ...users.User) {
	for _, user := range b {
		_ = s.rR.AddPlayer(r, user)
	}
}