package main

import (
	"fmt"
	"time"

	"github.com/bunterg/card-server/adding"
	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/storage"
)

func main() {
	start := time.Now()
	// set up storage
	storageType := storage.InMemory // this could be a flag; hardcoded here for simplicity
	var cardsStorage cards.Repository
	switch storageType {
	case storage.InMemory:
		cardsStorage = new(storage.MemoryCardStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
	}
	// create the available services
	adder := adding.NewService(cardsStorage)

	// add some sample data
	adder.AddSampleCards()

	// change rand seed
	fmt.Println(cardsStorage.GetAll()[105:])
	elapsed := time.Since(start)
	fmt.Println("scrypt took", elapsed)
}
