package authentication

import "net/http"

func UnauthorizedResponse(w http.ResponseWriter)  {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	w.Header().Set(
		"WWW-Authenticate",
		"API realm=Please login to use this site.",
	)
}
