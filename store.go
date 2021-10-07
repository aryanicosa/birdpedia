package main

import(
	"database/sql"
)

//create method
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds()([]*Bird, error)
}

//implement interface to struct
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES ($1, $2)", bird.Species, bird.Description)

	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {
	//query the database
	rows, err := store.db.Query("SELECT species, description FROM birds")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	birds := []*Bird{}

	for rows.Next() {
		//create pointer to a bird from each rows returned
		bird := &Bird{}

		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		//apend the result to returned array
		birds = append(birds, bird)
	}

	return birds, nil
}

var store Store

//init the store
func InitStore(s Store) {
	store = s
}