package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running (port=80), route: http://localhost:80/helloworld\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
