package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my portfolio")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Server was shot down")
	}
}
