package storage

import (
	"card-server/cards"
	"card-server/decks"
	"time"
)

// Memory storage keeps beer data in memory
type MemoryCardStorage struct {
	cards []cards.Card
}

// Add saves the given Card to the repository
func (m *MemoryCardStorage) Add(c cards.Card) error {
	// for _, e := range m.cards {
	// 	if c.Att == e.Att &&
	// 		c.Def == e.Def &&
	// 		c.Cost == e.Def &&
	// 		c.Class == e.Class {
	// 		return cards.ErrDuplicate
	// 	}
	// }

	c.ID = len(m.cards) + 1
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

// GetAll return all beers
func (m *MemoryCardStorage) GetAll() []cards.Card {
	return m.cards
}

// // Memory storage keeps review data in memory
// type MemoryReviewStorage struct {
// 	beers   []cards.Card
// 	reviews []reviews.Review
// }

// // Add saves the given review in the repository
// func (m *MemoryReviewStorage) Add(r reviews.Review) error {
// 	found := false
// 	for b := range m.beers {
// 		if m.beers[b].ID == r.BeerID {
// 			found = true
// 		}
// 	}

// 	if found {
// 		r.ID = fmt.Sprintf("%s_%s_%s_%s", r.BeerID, r.FirstName, r.LastName, r.Created.Unix())
// 		r.Created = time.Now()
// 		m.reviews = append(m.reviews, r)
// 	} else {
// 		return reviews.ErrNotFound
// 	}

// 	return nil
// }

// // GetAll returns all reviews for a given beer
// func (m *MemoryReviewStorage) GetAll(beerID int) []reviews.Review {
// 	var list []reviews.Review

// 	for i := range m.reviews {
// 		if m.reviews[i].BeerID == beerID {
// 			list = append(list, m.reviews[i])
// 		}
// 	}

// 	return list
// }

type MemoryDeckStorage struct {
	decks []decks.Deck
}

// Add saves the given Card to the repository
func (m *MemoryDeckStorage) Add(d decks.Deck) error {
	// for _, e := range m.cards {
	// 	if c.Att == e.Att &&
	// 		c.Def == e.Def &&
	// 		c.Cost == e.Def &&
	// 		c.Class == e.Class {
	// 		return cards.ErrDuplicate
	// 	}
	// }

	d.ID = len(m.decks) + 1
	d.Created = time.Now()
	m.decks = append(m.decks, d)

	return nil
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

// GetAll return all beers
func (m *MemoryDeckStorage) GetAll() []decks.Deck {
	return m.decks
}
