package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, AWS CI/CD with Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
