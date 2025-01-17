package data

import (
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"jervis/llm"
	"jervis/utils"
)

type Interaction struct {
	CurrentRequest   string
	CurrentAction    string
	PreviousRequest  string
	CurrentResponse  llm.Response `storm:"inline"`
	PreviousResponse llm.Response `storm:"inline"`
	Context          string
}

type MimsObject struct {
	ID          int `storm:"id,increment"`
	Key         string
	Summary     string
	Interaction `storm:"inline"`
}

var log = utils.Log

// Open a connection to the datavsse
func Open(path string) (*storm.DB, error) {
	db, err := storm.Open(path)
	if err != nil {
		log.Error("Error opening database:", err)
		return nil, err
	}

	return db, nil
}

// Put takes in a MimsObject and saves it to our db
func Put(message *MimsObject) error {
	
	// Open a database connection
	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return err
	}

	// Save our mims object
	err = db.Save(message)
	if err != nil {
		log.Error("Error saving data to db:", err)
		return err
	}

	// Close the db connection
	defer db.Close()
	
	return nil
}

// Search takes in a string and returns a list of matching mins objects
func Search(searchTerms string) ([]MimsObject, error) {
	
	// Open a DB connection
	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return messages, err
	}

	// Search the DB
  var results []MimsObject
	err = db.Select(q.Re("Summary", "^.*"+searchTerms+".*$")).Find(&results)
	if err != nil {
		log.Info("No results found for:", searchTerms)
		return results, err
	}
	
	// Close the db connextion
  defer db.Close()
	return results, nil
}
