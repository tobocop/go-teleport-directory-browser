package main

import (
	"github.com/tobocop/go-teleport-directory-browser/api"
	"log"
	"net/http"
)


func main() {
	api.RegisterApiRoutes()
	err := http.ListenAndServeTLS(":8080", "certs/localhost.crt", "certs/localhost.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

