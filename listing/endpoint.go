package listing

import (
	"encoding/json"
	"net/http"
)

// MakeGetUsersEndpoint generate user listing endpoing
func MakeGetUsersEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetUsers()
		json.NewEncoder(w).Encode(list)
	}
}

// MakeGetRoomsEndpoint generate rooms listing endpoing
func MakeGetRoomsEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetRooms()
		json.NewEncoder(w).Encode(list)
	}
}
