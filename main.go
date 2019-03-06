package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/bunterg/card-server/listing"
	"github.com/bunterg/card-server/rooms"

	"github.com/bunterg/card-server/adding"
	"github.com/bunterg/card-server/cards"
	"github.com/bunterg/card-server/storage"
	"github.com/bunterg/card-server/users"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("SERVE HOME")
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	// start := time.Now()
	// elapsed := time.Since(start)
	flag.Parse()

	// set up storage
	storageType := storage.InMemory // this could be a flag; hardcoded here for simplicity
	var cardsStorage cards.Repository
	var usersStorage users.Repository
	var roomsStorage rooms.Repository
	switch storageType {
	case storage.InMemory:
		cardsStorage = new(storage.MemoryCardStorage)
		usersStorage = new(storage.MemoryUserStorage)
		roomsStorage = new(storage.MemoryRoomStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
	}
	// create the available services
	cardAdder := cards.NewService(cardsStorage)
	adder := adding.NewService(usersStorage, roomsStorage)
	lister := listing.NewService(usersStorage, roomsStorage)

	// add default cards
	cardAdder.AddSampleCards()

	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/createUser/", adding.MakeAddUserEndpoint(adder))
	http.HandleFunc("/getUsers/", listing.MakeGetUsersEndpoint(lister))
	http.HandleFunc("/createRoom/", adding.MakeAddRoomEndpoint(adder))
	http.HandleFunc("/getRooms/", listing.MakeGetRoomsEndpoint(lister))
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		wsPath := "/ws/"
		if r.URL.Path == wsPath {
			http.Error(w, "No room params found", http.StatusNotFound)
			return
		}
		roomID := r.URL.Path[len(wsPath):]
		id, err := strconv.Atoi(roomID)
		if err != nil {
			http.Error(w, "Invalid room id", http.StatusBadRequest)
			return
		}
		_, err = roomsStorage.Get(id)
		if err != nil {
			http.Error(w, "Room not found", http.StatusNotFound)
			return
		}
		serveWs(hub, w, r, roomID)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
