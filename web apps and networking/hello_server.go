package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	myHandler := new(hello)
	http.Handle("/home", myHandler)
	http.ListenAndServe(":9000", myHandler)
}
