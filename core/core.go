package core

import (
	"fmt"
	"jervis/data"
	"strings"
)

func SearchHistory(i data.Interaction) ([]data.MimsObject, error) {
	updMessage := strings.ReplaceAll(i.CurrentRequest, "Search:", "")

	var SearchResult []data.MimsObject
	var err error
	SearchResult, err = data.Search(updMessage)
	if err != nil {
		fmt.Println("Error searching chats about: ", updMessage, err)
	}

	return SearchResult, nil
}

func SaveHistory(i data.MimsObject) error {

	if err := data.Put(&i); err != nil {
		fmt.Print("bad put")
	}

	return nil
}

func AddContext(i data.Interaction) {

}

func ResearchTopic(i data.Interaction) {

}
