package main

import (
	"card-server/cards"
	"card-server/decks"
	"card-server/storage"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// set up storage
	storageType := storage.InMemory // this could be a flag; hardcoded here for simplicity
	var cardsStorage cards.Repository
	var decksStorage decks.Repository
	switch storageType {
	case storage.InMemory:
		cardsStorage = new(storage.MemoryCardStorage)
		decksStorage = new(storage.MemoryDeckStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
		// cardsStorage, _ = storage.NewJSONBeerStorage()
		// decksStorage, _ = storage.NewJSONReviewStorage()
	}
	// change rand seed
	elapsed := time.Since(start)
	fmt.Println("scrypt took", elapsed)
}

// r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
// for index := 0; index < r.Intn(100); index++ {
// 	rand.Intn(10)
// }

// cardsID := rand.Perm(10)
// fmt.Println("O", cardsID)
// rand.Shuffle(len(cardsID), func(i, j int) {
// 	cardsID[i], cardsID[j] = cardsID[j], cardsID[i]
// })
// fmt.Println("C", cardsID)
// d := decks.Get()
// for _, c := range d.Cards {
// 	fmt.Println(c)
// }
