package main

import (
	"flag"
	"log"
	"net/http"

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
	switch storageType {
	case storage.InMemory:
		cardsStorage = new(storage.MemoryCardStorage)
		usersStorage = new(storage.MemoryUserStorage)
	case storage.JSONFiles:
		// error handling omitted for simplicity
	}
	// create the available services
	adder := cards.NewService(cardsStorage)
	userAdder := adding.NewService(usersStorage)
	// add some sample data
	adder.AddSampleCards()

	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/signup/", adding.MakeAddUserEndpoint(userAdder))
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		wsPath := "/ws/"
		if r.URL.Path == wsPath {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		room := r.URL.Path[len(wsPath):]
		serveWs(hub, w, r, room)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
