package main

import (
	"bufio"
	"fmt"
	"jervis/core"
	"jervis/data"
	"jervis/llm"
	"os"
	"strings"
)

func main() {

	// interaction object for this session
	var interaction data.Interaction

	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		// Get user input
		interaction.CurrentRequest = s.Text()

		// Print a cursor
		fmt.Print("> ")

		if strings.Contains(s.Text(), "Search:") {
			updMessage := strings.ReplaceAll(s.Text(), "Search:", "")

			searchObject := data.Interaction{
				CurrentRequest: updMessage,
			}

			searchResult, err := core.SearchHistory(searchObject)
			if err != nil {
				fmt.Println("Error searching chats about: ", updMessage, err)
			}
			fmt.Println("Here are the results: ")
			fmt.Println(searchResult)
			fmt.Println(">")
			continue
		}

		if strings.Contains(strings.ToLower(s.Text()), "research") {
			updMessage := strings.ReplaceAll(s.Text(), "research:", "")
			request := llm.Request{
				Query:   updMessage,
				Context: "",
			}

			var resp llm.Response
			resp, _ = llm.Query(request)
			interaction.CurrentResponse = resp
			fmt.Printf("Summary: %s\n\n", resp.Summary)
			fmt.Printf("Details: %s\n\n", resp.Details)
			fmt.Print("> ")
		}

		if strings.Contains(s.Text(), "End:") {
			break
		}
	}
}
