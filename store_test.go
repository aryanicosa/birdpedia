package main

import(
	"database/sql"
	//"testing"

	"fmt"
	// The "testify/suite" package is used to make the test suite
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite

	store *dbStore
	db *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	//open database connection and stored as an instance variable
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func (s *StoreSuite) TestCreateBird() {
	s.store.CreateBird(&Bird{
		Description: "test desc",
		Species: "test spec",
	})

	//query database
	res, err := s.db.Query("SELECT COUNT(*) FROM birds WHERE description='test desc' AND species='test spec'")
	if err != nil {
		s.T().Fatal(err)
	}

	//get count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1 but got %d", count)
	}
}

func (s *StoreSuite) TestGetBird() {
	//insert a sample to table
	_, err := s.db.Query("INSERT INTO birds (species, decsription) VALUES('bird', 'desc')")

	if err != nil {
		s.T().Fatal(err)
	}

	//get list
	birds, err := s.store.GetBirds()
	
	if err != nil {
		s.T().Fatal(err)
	}

	//assertion
	nBirds := len(birds)

	if nBirds != 1 {
		s.T().Errorf("incorrect count, wanted 1 but fot %d", nBirds)
	}

	expectedBird := Bird{"bird", "desc"}
	if *birds[0] != expectedBird {
		s.T().Errorf("incorrect details, expected %v but got %v", expectedBird, *birds[0])
	}

}