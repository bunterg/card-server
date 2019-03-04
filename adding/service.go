package adding

import "github.com/bunterg/card-server/users"

// Service provides User adding operations
type Service interface {
	AddUser(...users.User)
}

type service struct {
	uR users.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(uR users.Repository) Service {
	return &service{uR}
}

// AddUser adds the given user(s) to the database
func (s *service) AddUser(u ...users.User) {
	for _, user := range u {
		_ = s.uR.Add(user)
	}
}
