package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello I'm Hannah Mae\n")
	})

	fmt.Println("Hannah Mae's server.")
	http.ListenAndServe(":8080", nil)
}
