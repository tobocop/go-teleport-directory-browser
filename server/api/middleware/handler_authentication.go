package middleware

import (
	"github.com/tobocop/go-teleport-directory-browser/api/session"
	"net/http"
)

type handlerAuthenticator struct {
	sessionManager session.Manager
}

func NewHandlerAuthenticator(sessionManager session.Manager) *handlerAuthenticator {
	return &handlerAuthenticator{sessionManager}
}

func (s *handlerAuthenticator) AuthenticatedHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set(
				"WWW-Authenticate",
				"API realm=Please enter a valid username and password to use this site.",
			)
			return
		}

		err = s.sessionManager.ValidateSession(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set(
				"WWW-Authenticate",
				"API realm=Please enter a valid username and password to use this site.",
			)
			return
		}


		next.ServeHTTP(w, r)
	}
}