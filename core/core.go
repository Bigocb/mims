package core

import (
	"fmt"
	"jervis/data"
	"strings"
)

func ProcessUserInput(i data.Interaction) {
	if strings.Contains(i.CurrentRequest, "Search:") {
		updMessage := strings.ReplaceAll(i.CurrentRequest, "Search:", "")

		var SearchResult data.Interaction
		var err error
		SearchResult, err = data.Search(updMessage)
		if err != nil {
			fmt.Println("Error searching chats about: ", updMessage, err)
		}
		fmt.Println("Here are the results: ")
		fmt.Println(SearchResult)
		fmt.Println(">")
	}
}
