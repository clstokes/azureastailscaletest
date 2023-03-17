package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
		fmt.Fprintf(w, "\n")

		envs := os.Environ()
		sort.Strings(envs)
		for _, e := range envs {
			// pair := strings.SplitN(e, "=", 2)
			// fmt.Fprintf(w, pair[0])
			fmt.Fprintf(w, e)
			fmt.Fprintf(w, "\n")
		}

	})
	fmt.Printf("Server running (port=80), route: http://localhost:80/helloworld\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
