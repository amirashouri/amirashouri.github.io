package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my portfolio")
	})

	listenAddr := os.Getenv("LISTEN_ADDR")
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic("Server was shot down")
	}
}
