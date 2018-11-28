package main

import "math/rand"

func getAll() []Card {
	var cardLibrary []Card
	for index := 0; index < 50; index++ {
		cardLibrary = append(
			cardLibrary,
			Card{index, rand.Intn(5), rand.Intn(5), rand.Intn(5), CardType(rand.Intn(3))})
	}

	return cardLibrary
}
