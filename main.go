//everything with package name can see eeverything else inside the same package
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" //import gorilla/mux library
)

func newRouter() *mux.Router {
	//declare a new router
	r := mux.NewRouter() 

	r.HandleFunc("/hello", handler).Methods("GET")

	//declare static file directory
	staticFileDirectory := http.Dir("./assets")

	//delete "/assets" prefix and file handler
	staticFileHandler := http.StripPrefix("/assets", http.FileServer(staticFileDirectory))

	//create matcher for all routes starting with "/assets/" and give the handler and method
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	r := newRouter()

	//http.HandleFunc("/", handler) // it's us basic http router

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo dunia!")
}