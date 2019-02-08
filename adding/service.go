package adding

import (
	"github.com/bunterg/card-server/cards"
)

// Service provides card adding operations
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

// AddCard adds the given card(s) to the database
func (s *service) AddCard(b ...cards.Card) {
	for _, card := range b {
		_ = s.cR.Add(card) // error handling omitted for simplicity
	}
}

// AddSampleCards adds some sample cards to the database
func (s *service) AddSampleCards() {
	for _, b := range cards.DefaultCards {
		_ = s.cR.Add(b) // error handling omitted for simplicity
	}
}
