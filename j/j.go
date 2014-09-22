package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Hi there, I love %s! \n", r.URL.Path[1:])
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Web server running: localhost:8080")

    http.ListenAndServe(":8080", nil)
}