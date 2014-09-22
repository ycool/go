package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (v String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,string(v))
}

func (v *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, v.Greeting, v.Punct, v.Who)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    
    http.ListenAndServe("localhost:4000", nil)
}