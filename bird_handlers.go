package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species string "json:\"species\""
	Description string "json:\"description\""
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	
	birds, err := store.GetBirds()
	
	//convert bird to json
	birdListBytes, err := json.Marshal(birds)


	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	//create bird instance
	bird := Bird{}

	//form parser
	err := r.ParseForm()

	//form parser error handler
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//get information from the form
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	//append data to existing list with a new entry
	//birds = append(birds, bird)

	err = store.CreateBird(&bird)
	
	if err != nil {
		fmt.Println(err)
	}

	//redirect user to the original html page
	http.Redirect(w, r, "/assets/", http.StatusFound)
}