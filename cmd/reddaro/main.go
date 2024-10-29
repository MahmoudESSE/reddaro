package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var posts = []post{
	{ID: "1", Title: "Where is waldo?"},
	{ID: "2", Title: "Who is waldo?"},
	{ID: "3", Title: "How is waldo?"},
}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	fmt.Fprintf(w, "Hello, you requested: %s\n", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

// Rest HTTP GET request to get back all post in the database
func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /posts request\n")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getClickedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /clicked request\n")
	fmt.Fprintf(w, "I was clicked!")

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api", getRootHandler)
	http.HandleFunc("/api/posts", getPostsHandler)
	http.HandleFunc("/api/clicked", getClickedHandler)

	fs := http.FileServer(http.Dir("web/"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
