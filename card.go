package main

// CardType indicate type of card
type CardType int

// Card data model for game card profile
type Card struct {
	id    int
	att   int
	def   int
	cost  int
	class CardType
}

const (
	basic CardType = iota
	special
	legendary
)

func (cardType CardType) String() string {
	names := [...]string{
		"Basic",
		"Special",
		"Legendary"}

	// invalid cardType, out of range
	if cardType < basic || cardType > legendary {
		return "Unknown"
	}

	return names[cardType]
}
