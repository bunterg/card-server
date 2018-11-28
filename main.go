package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	// change rand seed
	// r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	// for index := 0; index < r.Intn(100); index++ {
	// 	rand.Intn(10)
	// }

	cardsID := rand.Perm(10)
	fmt.Println("O", cardsID)
	rand.Shuffle(len(cardsID), func(i, j int) {
		cardsID[i], cardsID[j] = cardsID[j], cardsID[i]
	})
	fmt.Println("C", cardsID)

	elapsed := time.Since(start)
	fmt.Println("scrypt took", elapsed)
}
