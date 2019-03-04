package adding

import (
	"github.com/bunterg/card-server/rooms"
	"github.com/bunterg/card-server/users"
)

// Service provides User adding operations
type Service interface {
	AddUser(...users.User) []users.User
	AddRoom(u users.User) rooms.Room
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
func (s *service) AddUser(u ...users.User) []users.User {
	var ru []users.User
	for _, user := range u {
		nu, _ := s.uR.Add(user)
		ru = append(ru, nu)
	}
	return ru
}

func (s *service) AddRoom(u users.User) rooms.Room {
	r, _ := s.rR.Add(rooms.Room{}, u)
	return r
}
