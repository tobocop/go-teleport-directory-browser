package api

import (
	"github.com/tobocop/go-teleport-directory-browser/api/authentication"
	"net/http"
)

func RegisterApiRoutes(){
	authServer := authentication.NewServer()
	http.HandleFunc("/api/authenticate", authServer.AuthHandler)
}
