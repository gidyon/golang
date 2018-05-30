package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	mux2 := http.NewServeMux()
	fb := FB{"max", "kirwa", "simon", "david"}
	mux2.Handle("/home", http.HandlerFunc(fb.home))
	mux2.Handle("/friends", http.HandlerFunc(fb.friends))
	go startDBServer(mux)
	log.Fatal(http.ListenAndServe("localhost:8000", mux2))
}
func startDBServer(mux http.Handler) {
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type FB []string

func (f FB) home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Friends")
}
func (f FB) friends(w http.ResponseWriter, req *http.Request) {
	for _, friends := range f {
		fmt.Fprintf(w, friends)
	}
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
