package adding

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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
		// SUCCESS
		w.WriteHeader(http.StatusOK)
		s.AddUser(users.User{Name: m.Username})
		io.WriteString(w, "New user: "+m.Username)
	}
}
