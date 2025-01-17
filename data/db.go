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

func Open(path string) (*storm.DB, error) {
	db, err := storm.Open(path)
	if err != nil {
		log.Error("Error opening database:", err)
		return nil, err
	}

	return db, nil
}

func Put(message *MimsObject) error {
	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return err
	}

	err = db.Save(message)
	if err != nil {
		log.Error("Error saving data to db:", err)
		return err
	}

	var users []MimsObject
	err = db.All(&users)
	if err != nil {
		fmt.Print("test")
	}

	defer db.Close()
	return nil
}

func Search(searchTerms string) ([]MimsObject, error) {
	var messages []MimsObject
	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return messages, err
	}

	var users []MimsObject
	err = db.All(&users)

	err = db.Select(q.Re("Summary", "^.*"+searchTerms+".*$")).Find(&messages)
	if err != nil {
		log.Info("No results found for:", searchTerms)
		return messages, err
	}

	return messages, nil
}
