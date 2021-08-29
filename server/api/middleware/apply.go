package middleware

import "net/http"

func Apply(method string, next http.HandlerFunc) http.HandlerFunc  {
	return csrf(methodNotAllowed(method, next))
}
