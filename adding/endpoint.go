package adding

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bunterg/card-server/users"
)

// SignUpBody body for signup request
type SignUpBody struct {
	Username string `json:"username"`
}

// MakeAddUserEndpoint user endpoint
func MakeAddUserEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SIGNUP FUNC")
		if r.URL.Path != "/signup/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request (body)", http.StatusBadRequest)
			return
		}
		// var m map[string]interface{}
		var m SignUpBody
		err = json.Unmarshal(body, &m)
		if err != nil {
			http.Error(w, "Bad Request (body data)", http.StatusBadRequest)
			return
		}
		log.Println("NEW USER: " + m.Username)
		profile := s.AddUser(users.User{Name: m.Username})
		js, err := json.Marshal(profile[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// SUCCESS
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

// NewRoomBody body for room create request
type NewRoomBody struct {
	Owner users.User `json:"owner"`
}

// MakeAddRoomEndpoint Room endpoint
func MakeAddRoomEndpoint(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("NEW ROOM FUNC")
		if r.URL.Path != "/createRoom/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request (body)", http.StatusBadRequest)
			return
		}
		// var m map[string]interface{}
		var m NewRoomBody
		err = json.Unmarshal(body, &m)
		if err != nil {
			http.Error(w, "Bad Request (body data)", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		room := s.AddRoom(m.Owner)
		log.Println("NEW ROOM: ")
		io.WriteString(w, "New user: "+strconv.Itoa(room.ID))
	}
}
