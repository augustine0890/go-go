package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":8080", nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name)
}
