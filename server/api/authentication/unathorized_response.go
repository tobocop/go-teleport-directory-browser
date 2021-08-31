package authentication

import "net/http"

func UnauthorizedResponse(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set(
		"WWW-Authenticate",
		"API realm=Please login to use this site.",
	)
}
