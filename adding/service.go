package adding

import (
	"github.com/bunterg/card-server/cards"
)

// Service provides beer adding operations
type Service interface {
	AddCard(...cards.Card)
	AddSampleCards()
}

type service struct {
	cR cards.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(cR cards.Repository) Service {
	return &service{cR}
}

// AddCard adds the given beer(s) to the database
func (s *service) AddCard(b ...cards.Card) {
	for _, beer := range b {
		_ = s.cR.Add(beer) // error handling omitted for simplicity
	}
}

// AddSampleCards adds some sample beers to the database
func (s *service) AddSampleCards() {
	for _, b := range cards.DefaultCards {
		_ = s.cR.Add(b) // error handling omitted for simplicity
	}
}
