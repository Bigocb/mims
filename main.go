package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jervis/core"
	"jervis/data"
	"jervis/llm"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// interaction object for this session
	var interaction core.Interaction

	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		// Get user input
		interaction.CurrentRequest = s.Text()

		// Print a cursor
		fmt.Print("> ")

		// Process user input.
		core.ProcessUserInput(interaction)

		if strings.Contains(s.Text(), "Add Context:") {
			out, err := json.Marshal(interaction.SearchResult)
			if err != nil {
				fmt.Println("Error marshalling search result:", err)
			}

			interaction.Context = "These are previous responses on this topic. Please use these as context when responding. " + string(out)
			fmt.Println(interaction.Context)
		}

		if strings.Contains(strings.ToLower(s.Text()), "research") {
			request := llm.Request{
				Query:   interaction.CurrentRequest,
				Context: interaction.Context,
			}

			var resp llm.Response
			resp = llm.Query(request)
			interaction.PreviousResponse = resp
			fmt.Printf("Summary: %s\n\n", resp.Summary)
			fmt.Printf("Details: %s\n\n", resp.Details)
			fmt.Print("> ")
		}
		if strings.Contains(s.Text(), "End:") {
			break
		}
		if strings.Contains(s.Text(), "Save:") {
			timeNow := time.Now().Unix()

			messageObject := data.MimsData{
				Key:     strconv.FormatInt(timeNow, 10),
				Summary: interaction.PreviousResponse.Summary,
				Details: interaction.PreviousResponse.Details,
			}

			if err := data.Put(&messageObject); err != nil {
				fmt.Print("bad put")
			}
		}

	}
}
