package main

import (
	"context"
	"github.com/tobocop/go-teleport-directory-browser/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		osCall := <-c
		log.Printf("system call: %+v", osCall)
		cancel()
	}()

	err := serve(ctx)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}

func serve(ctx context.Context) (err error) {
	api.RegisterApiRoutes()
	srv := &http.Server{Addr: ":8080"}

	go func() {
		err := srv.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key")
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Println("Server started on port 8080")

	<-ctx.Done()

	log.Println("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
