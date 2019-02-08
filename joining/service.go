package joining

import (
	"github.com/bunterg/card-server/rooms"
	"github.com/bunterg/card-server/users"
)

// Service provides beer adding operations
type Service interface {
	AddPlayer(...users.User)
}

type service struct {
	rR rooms.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(rR rooms.Repository) Service {
	return &service{rR}
}

// AddCard adds the given user(s) to the database
func (s *service) AddPlayer(b ...users.User) {
	for _, user := range b {
		_ = s.rR.AddPlayer(user)
	}
}
