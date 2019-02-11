package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/storage"
	"github.com/bunterg/card-server/users"
)

func main() {
	start := time.Now()
	// set up storage
	storageType := storage.InMemory // this could be a flag; hardcoded here for simplicity
	var cardsStorage cards.Repository
	var usersStorage users.Repository
	switch storageType {
	case storage.InMemory:
		cardsStorage = new(storage.MemoryCardStorage)
		usersStorage = new(storage.MemoryUserStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
	}
	// create the available services
	adder := cards.NewService(cardsStorage)

	// add some sample data
	adder.AddSampleCards()
	usersStorage.Add(users.User{
		ID:      1,
		Name:    "Bernardo Garcia",
		Created: time.Now()})
	// change rand seed
	fmt.Println(cardsStorage.GetAll()[105:])
	fmt.Println(usersStorage.GetAll())
	elapsed := time.Since(start)
	fmt.Println("scrypt took", elapsed)
	os.Exit(1)
}
