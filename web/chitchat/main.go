package chitchat

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("ChitChat %s started at %s", version(), config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
