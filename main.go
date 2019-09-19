package main

import (
	"fmt"
	"log"
	"net/http"
)

const resp = "Hi (not novian here), my name is %s"

func main() {
	http.HandleFunc("/", greetingHandler)
	log.Fatal(http.ListenAndServe(":8989", nil))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, resp, r.URL.Path[1:])
}
