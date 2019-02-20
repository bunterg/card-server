package decks

import "github.com/bunterg/card-server/cards"

// Service provides card adding operations
type Service interface {
	AddDeck(...Deck)
}

type service struct {
	dR Repository
	cR cards.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(dR Repository, cR cards.Repository) Service {
	return &service{dR, cR}
}

// AddDeck adds the given card(s) to the database
func (s *service) AddDeck(ds ...Deck) {
	for _, deck := range ds {
		s.dR.Add(deck)
	}
}
