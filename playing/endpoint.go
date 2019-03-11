package playing

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bunterg/card-server/users"
)

// MakeGameRoomEndpoint user endpoint
func MakeGameRoomEndpoint(s Service, hub *Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		// check if room exist
		// todo check room is open
		_, err = s.GetRoom(id)
		if err != nil {
			http.Error(w, "Room not found", http.StatusNotFound)
			return
		}
		// Identify user
		keys := r.URL.Query()
		var user users.User
		err = json.Unmarshal([]byte(keys["user"][0]), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// check if room exist
		// todo check room is open
		_, err = s.GetUser(user.ID)
		serveWs(hub, w, r, r.URL.Path[len(wsPath):])
	}
}
