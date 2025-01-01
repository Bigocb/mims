package data

import (
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
)

type JervisData struct {
	ID      int `storm:"id,increment"`
	Key     string
	Message string `storm:"index"`
}

func Open(path string) (*storm.DB, error) {
	db, err := storm.Open(path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Put(message *JervisData) error {
	db, err := Open("content.db")
	if err != nil {
		return err
	}

	err = db.Init(JervisData{})
	if err != nil {
		return err
	}

	err = db.Save(message)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

func Search(searchTerms string) ([]JervisData, error) {
	db, err := Open("content.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Searching ", searchTerms)
	var messages []JervisData
	err = db.Select(q.Re("Message", searchTerms)).Find(&messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
