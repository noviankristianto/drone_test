package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", greetingHandler)
	log.Fatal(http.ListenAndServe(":8989", nil))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	// test
	fmt.Fprintf(w, "Hi there, my name is %s", r.URL.Path[1:])
}
