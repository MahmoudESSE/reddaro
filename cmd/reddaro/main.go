package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type post struct {
	ID    string
	Title string
}

var posts = []post{
	{ID: "1", Title: "Where is waldo?"},
	{ID: "2", Title: "Who is waldo?"},
	{ID: "3", Title: "How is waldo?"},
}

type api struct{}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello, World!\n")
}

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

func main() {
	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/posts", getPostsHandler)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
