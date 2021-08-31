package api

import (
	"fmt"
	"github.com/tobocop/go-teleport-directory-browser/api/authentication"
	"github.com/tobocop/go-teleport-directory-browser/api/middleware"
	"github.com/tobocop/go-teleport-directory-browser/api/session"
	"net/http"
)

type route struct {
	path          string
	method        string
	authenticated bool
	handler       http.HandlerFunc
}

func RegisterApiRoutes() {
	sessionManager := session.NewInMemoryManager()
	auth := middleware.NewHandlerAuthenticator(sessionManager)
	authServer := authentication.NewServer(sessionManager)

	routes := []route{
		{"/authenticate", http.MethodPost, false, authServer.AuthHandler},
		{"/me", http.MethodGet, true, authServer.Me},
	}

	for _, r := range routes {
		h := middleware.Apply(r.method, r.handler)
		if r.authenticated {
			h = auth.AuthenticatedHandler(h)
		}
		http.HandleFunc( fmt.Sprintf("/api%s", r.path), h)
	}
}
