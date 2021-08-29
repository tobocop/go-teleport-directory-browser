package api

import (
	"github.com/tobocop/go-teleport-directory-browser/api/authentication"
	"github.com/tobocop/go-teleport-directory-browser/api/middleware"
	"net/http"
)

func RegisterApiRoutes(){
	authServer := authentication.NewServer()
	http.HandleFunc("/api/authenticate", middleware.Apply(http.MethodPost, authServer.AuthHandler))
}
