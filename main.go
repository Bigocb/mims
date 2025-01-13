package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jervis/data"
	"jervis/llm"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Print("> ")

	s := bufio.NewScanner(os.Stdin)
	var searchResult []data.JervisData
	var err error
	for s.Scan() {

		requestText := s.Text()
		fmt.Print("> ")

		if strings.Contains(s.Text(), "Search:") {
			updMessage := strings.ReplaceAll(s.Text(), "Search:", "")

			searchResult, err = data.Search(updMessage)
			if err != nil {
				fmt.Println("Error searching chats about:", err)
			}
			fmt.Println("Here are the results: ")
			fmt.Println(searchResult)
			fmt.Println(">")
			continue
		}

		if strings.Contains(s.Text(), "Add Context:") {
			out, err := json.Marshal(searchResult)
			if err != nil {
				fmt.Println("Error marshalling search result:", err)
			}
			var updateText string
			updateText = "These are previous responses on this topic. Please use these as context when responding. " + string(out)
			requestText = updateText
		}

		if strings.Contains(strings.ToLower(s.Text()), "research") {
			var resp llm.Response
			resp = llm.Query(requestText)
			fmt.Printf("Summary: %s\n\n", resp.Summary)
			fmt.Printf("Details: %s\n\n", resp.Details)
			fmt.Print("> ")
		}

		if strings.Contains(s.Text(), "End:") {
			break
		}
		if strings.Contains(s.Text(), "Save:") {
			timeNow := time.Now().Unix()

			messageObject := data.JervisData{
				Key:     strconv.FormatInt(timeNow, 10),
				Message: requestText,
			}

			if err := data.Put(&messageObject); err != nil {
				fmt.Print("bad put")
			}
		}

	}
}
