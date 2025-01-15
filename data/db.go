package data

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"jervis/utils"
)

var log = utils.Log

type MimsData struct {
	ID      int `storm:"id,increment"`
	Key     string
	Details string `storm:"index"`
	Summary string `storm:"index"`
}

func Open(path string) (*storm.DB, error) {
	db, err := storm.Open(path)
	if err != nil {
		log.Error("Error opening database:", err)
		return nil, err
	}

	return db, nil
}

func Put(message *MimsData) error {
	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return err
	}

	err = db.Init(MimsData{})
	if err != nil {
		log.Error("Error opening database:", err)
		return err
	}

	err = db.Save(message)
	if err != nil {
		log.Error("Error saving data to db:", err)
		return err
	}

	defer db.Close()
	return nil
}

func Search(searchTerms string) ([]MimsData, error) {

	db, err := Open("content.db")
	if err != nil {
		log.Error("Error opening database:", err)
		return nil, err
	}

	var messages []MimsData
	err = db.Select(q.Re("Details", "^.*"+searchTerms+".*$")).Find(&messages)
	if err != nil {
		log.Info("No results found for:", searchTerms)
		return nil, err
	}

	return messages, nil
}
