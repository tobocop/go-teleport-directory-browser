package middleware

import "net/http"

func csrf(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("csrf-token")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("X-CSRF-Token") != cookie.Value {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
