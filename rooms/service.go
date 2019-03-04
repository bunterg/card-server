package rooms

import (
	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/decks"
	"github.com/bunterg/card-server/users"
)

// Service provides player adding operations
type Service interface {
	AddPlayer(Room, ...users.User)
	NewRoom(users.User) Room
}

type service struct {
	rR Repository
	dR decks.Repository
	cR cards.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(rR Repository, dR decks.Repository, cR cards.Repository) Service {
	return &service{rR, dR, cR}
}

func (s *service) NewRoom(u users.User) Room {
	us := []users.User{u}
	r, _ := s.rR.Add(Room{
		Players: us,
		Owner:   u,
	}, u)
	return r
}

// AddCard adds the given user(s) to the database
func (s *service) AddPlayer(r Room, b ...users.User) {
	for _, user := range b {
		_ = s.rR.AddPlayer(r, user)
	}
}

// InitMatch begins room match
func (s *service) InitMatch(r Room) error {
	if len(r.Players) < 2 {
		return ErrNotEnoughPlayers
	}
	/// TODO check if game has begun
	// if r.Board == Board{} {
	// 	return
	// }
	Turn := 0
	Graveyard, _ := s.dR.Add(decks.Deck{})
	Deck, _ := s.dR.Add(decks.Deck{
		Cards: s.cR.GetAll(),
	})
	Deck.Shuffle()
	var Hands []Hand
	for _, player := range r.Players {
		cards, _ := Deck.Draw(7)
		Hands = append(Hands, Hand{
			Cards:  cards,
			Player: player,
		})
	}
	r.Board = Board{
		Graveyard,
		Deck,
		Turn,
		Hands,
	}
	return nil
}
