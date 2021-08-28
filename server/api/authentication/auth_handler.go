package authentication

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) AuthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var authReq AuthRequest
	err := json.NewDecoder(req.Body).Decode(&authReq)
	if err != nil {
		log.Printf("AuthHandler json decode error: %v", err)
		w.WriteHeader(400)
	}

	escapedUser, err := url.QueryUnescape(authReq.Username)
	if err != nil {
		log.Printf("AuthHandler username escape error: %v", err)
		w.WriteHeader(500)
		return
	}

	escapedPass, err := url.QueryUnescape(authReq.Password)
	if err != nil {
		log.Printf("AuthHandler password escape error: %v", err)
		w.WriteHeader(500)
		return
	}

	if s.Authenticator.Authenticate(escapedUser, escapedPass) {
		w.WriteHeader(204)
	} else {
		w.Header().Set(
			"WWW-Authenticate",
			"API realm=Please enter a valid username and password to use this site.",
		)
		w.WriteHeader(401)
	}
}
