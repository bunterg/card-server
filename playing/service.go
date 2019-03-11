package playing

import (
	"github.com/bunterg/card-server/rooms"
	"github.com/bunterg/card-server/users"
)

// Service provides player adding operations
type Service interface {
	GetUsers() []users.User
	GetRooms() []rooms.Room
	GetUser(id int) (users.User, error)
	GetRoom(id int) (rooms.Room, error)
}

type service struct {
	uR users.Repository
	rR rooms.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(uR users.Repository, rR rooms.Repository) Service {
	return &service{uR, rR}
}

func (s *service) GetUsers() []users.User {
	return s.uR.GetAll()
}

func (s *service) GetRooms() []rooms.Room {
	return s.rR.GetAll()
}

func (s *service) GetUser(id int) (users.User, error) {
	return users.User{}, nil
}

func (s *service) GetRoom(id int) (rooms.Room, error) {
	return rooms.Room{}, nil
}
