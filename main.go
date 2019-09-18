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
	fmt.Fprintf(w, "Hi there, my name is %s, nice to meet you!", r.URL.Path[1:])
}
