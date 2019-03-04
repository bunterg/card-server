package adding

import (
	"github.com/bunterg/card-server/rooms"
	"github.com/bunterg/card-server/users"
)

// Service provides User adding operations
type Service interface {
	AddUser(users.User) (users.User, error)
	AddRoom(users.User) (rooms.Room, error)
}

type service struct {
	uR users.Repository
	rR rooms.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(uR users.Repository, rR rooms.Repository) Service {
	return &service{uR, rR}
}

// AddUser adds the given user(s) to the database
func (s *service) AddUser(u users.User) (users.User, error) {
	return s.uR.Add(u)
}

func (s *service) AddRoom(u users.User) (rooms.Room, error) {
	user, _ := s.uR.Get(u.ID)
	return s.rR.Add(rooms.Room{}, user)
}
