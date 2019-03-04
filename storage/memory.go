package storage

import (
	"time"

	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/decks"
	"github.com/bunterg/card-server/rooms"
	"github.com/bunterg/card-server/users"
)

// ----------------CARDS-----------
// MemoryCardStorage storage keeps card data in memory
type MemoryCardStorage struct {
	cards []cards.Card
}

// Add saves the given Card to the repository
func (m *MemoryCardStorage) Add(c cards.Card) error {
	c.ID = len(m.cards)
	c.Created = time.Now()
	m.cards = append(m.cards, c)

	return nil
}

// Get returns a card with the specified ID
func (m *MemoryCardStorage) Get(id int) (cards.Card, error) {
	var card cards.Card

	for i := range m.cards {

		if m.cards[i].ID == id {
			return m.cards[i], nil
		}
	}

	return card, cards.ErrNotFound
}

// GetAll return all Cards
func (m *MemoryCardStorage) GetAll() []cards.Card {
	return m.cards
}

// ----------------DECK-----------

// MemoryDeckStorage deck storage}
// if r.Board == Board{} {
// 	return
// }}
// if r.Board == Board{} {
// 	return
// }
type MemoryDeckStorage struct {
	decks []decks.Deck
}

// Add saves the given Card to the repository
func (m *MemoryDeckStorage) Add(d decks.Deck) (decks.Deck, error) {
	d.ID = len(m.decks) + 1
	d.Created = time.Now()
	m.decks = append(m.decks, d)

	return d, nil
}

// Get returns a card with the specified ID
func (m *MemoryDeckStorage) Get(id int) (decks.Deck, error) {
	var d decks.Deck

	for i := range m.decks {

		if m.decks[i].ID == id {
			return m.decks[i], nil
		}
	}

	return d, decks.ErrNotFound
}

// GetAll return all decks
func (m *MemoryDeckStorage) GetAll() []decks.Deck {
	return m.decks
}

// ----------------USERS-----------
type MemoryUserStorage struct {
	users []users.User
}

// Add saves the given User to the repository
func (m *MemoryUserStorage) Add(u users.User) (users.User, error) {
	u.ID = len(m.users) + 1
	u.Created = time.Now()
	m.users = append(m.users, u)

	return u, nil
}

// Get returns a User with the specified ID
func (m *MemoryUserStorage) Get(id int) (users.User, error) {
	var d users.User

	for i := range m.users {

		if m.users[i].ID == id {
			return m.users[i], nil
		}
	}

	return d, users.ErrNotFound
}

// GetAll return all Users
func (m *MemoryUserStorage) GetAll() []users.User {
	return m.users
}

// ----------------DECK-----------

// MemoryRoomStorage idk yet
type MemoryRoomStorage struct {
	rooms []rooms.Room
}

// Add saves the given Card to the repository
func (m *MemoryRoomStorage) Add(r rooms.Room, u users.User) (rooms.Room, error) {
	r.ID = len(m.rooms)
	r.Created = time.Now()
	r.Owner = u
	r.Players = []users.User{u}
	m.rooms = append(m.rooms, r)
	return r, nil
}

// AddPlayer appends player to room
func (m *MemoryRoomStorage) AddPlayer(room rooms.Room, u users.User) error {
	r, _ := m.Get(room.ID)
	r.Players = append(r.Players, u)
	return nil
}

// Get returns a card with the specified ID
func (m *MemoryRoomStorage) Get(id int) (rooms.Room, error) {
	var d rooms.Room
	for i := range m.rooms {

		if m.rooms[i].ID == id {
			return m.rooms[i], nil
		}
	}
	return d, rooms.ErrNotFound
}

// GetAll return all rooms
func (m *MemoryRoomStorage) GetAll() []rooms.Room {
	return m.rooms
}
