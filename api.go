package main

import (
	"fmt"
	"net/http"
	"os"
)

var count int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request", count)
	count++
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	count = 1
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}