package adding

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bunterg/card-server/users"
)

// MakeAddUserEndpoint user endpoint
func MakeAddUserEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var m users.User
		err := postRequestData(w, r, "/createUser/", &m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		profile, _ := s.AddUser(m)
		log.Printf("NEW USER:\n %v", profile)
		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// SUCCESS
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

// MakeAddRoomEndpoint Room endpoint
func MakeAddRoomEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var m users.User
		err := postRequestData(w, r, "/createRoom/", &m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		room, _ := s.AddRoom(m)
		log.Printf("NEW ROOM:\n %v", room)
		js, err := json.Marshal(room)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// SUCCESS
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

func postRequestData(w http.ResponseWriter, r *http.Request, path string, m interface{}) error {
	if r.URL.Path != path {
		http.Error(w, "Not found", http.StatusNotFound)
		return errors.New("Not found")
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return errors.New("Method not allowed")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request (body)", http.StatusBadRequest)
		return errors.New("Bad Request (body)")
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		http.Error(w, "Bad Request (body data)", http.StatusBadRequest)
		return errors.New("Bad Request (body data)")
	}
	return nil
}
