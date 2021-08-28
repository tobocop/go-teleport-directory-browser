package main

import (
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from https.\n"))
}

func main() {
	http.HandleFunc("/api/hello", HelloServer)
	err := http.ListenAndServeTLS(":8080", "certs/localhost.crt", "certs/localhost.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

