package authentication

import (
	"encoding/json"
	"github.com/tobocop/go-teleport-directory-browser/api/session"
	"log"
	"net/http"
	"net/url"
	"time"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Not sure if query escape is best option. Maybe base64?
	escapedUser, err := url.QueryUnescape(authReq.Username)
	if err != nil {
		log.Printf("AuthHandler username escape error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	escapedPass, err := url.QueryUnescape(authReq.Password)
	if err != nil {
		log.Printf("AuthHandler password escape error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authenticated, err := s.Authenticator.Authenticate(escapedUser, escapedPass)
	if err != nil {
		log.Printf("AuthHandler authenticator error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if authenticated {
		sessionId, err := s.SessionManager.NewSession()
		if err != nil {
			log.Printf("AuthHandler session manager error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.ExpiresIn),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set(
			"WWW-Authenticate",
			"API realm=Please enter a valid username and password to use this site.",
		)
		w.WriteHeader(http.StatusUnauthorized)
	}
}
