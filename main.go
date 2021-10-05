//everything with package name can see eeverything else inside the same package
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" //import gorilla/mux library
)

func main() {
	//declare a new router
	r := mux.NewRouter() 

	r.HandleFunc("/hello", handler).Methods("GET")

	//http.HandleFunc("/", handler) // it's us basic http router

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo dunia!")
}