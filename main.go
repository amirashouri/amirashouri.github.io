package main

import (
	"log"
	"net/http"
	"os"
	"portfolio/handlers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.Make(handlers.HandleHome))

	listenAddr := os.Getenv("LISTEN_ADDR")
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic("Server was shot down")
	}
}
