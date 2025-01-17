package main

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v3"
	"jervis/data"
	"jervis/llm"
	"log"
	"os"
	"strings"
)

// interaction object for a given session
var interaction data.Interaction

// QueryLLM takes in a query string and sends to the LLM for processing.
func QueryLLM(question string) (string, error) {
	request := llm.Request{
		Query: question,
	}
	var resp llm.Response
	resp, err := llm.Query(request)
	if err != nil {
		return "", fmt.Errorf("error querying LLM: %w", err)
	}
	return resp.Details, nil
}

// QueryHistory takes in a query string and search the bolt db for related saves
func find(ctx context.Context, c *cli.Command) error {
	updMessage := strings.ReplaceAll("question", "Search:", "")
	var SearchResult data.Interaction
	var err error
	SearchResult, err = data.Search(os.Args[3])
	if err != nil {
		fmt.Println("Error searching chats about: ", updMessage, err)
	}
	fmt.Println("Here are the results: " + SearchResult.PreviousResponse.Details)
	return nil
}

func main() {

	cmd := &cli.Command{
		Name:  "mims",
		Usage: "Get response from AI",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "ponder",
				Aliases: []string{"p", "pond"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "topic",
						Aliases:     []string{"top", "t"},
						Value:       "",
						Usage:       "Please enter a topic",
						Destination: &interaction.CurrentRequest,
					},
				},
				Category: "LLM",
				Action:   ponder,
			},
			&cli.Command{
				Name:    "keep",
				Aliases: []string{"k"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "user",
						Aliases:     []string{"u"},
						Value:       "",
						Usage:       "Enter what you want to keep",
						Destination: &interaction.CurrentRequest,
					},
					&cli.StringFlag{
						Name:        "last",
						Aliases:     []string{"l"},
						Value:       "",
						Usage:       "Store the previous response for good measure",
						Destination: &interaction.CurrentRequest,
					},
				},
				Category: "Storage",
				Action:   keep,
			},
			&cli.Command{
				Name:    "find",
				Aliases: []string{""},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "topic",
						Aliases:     []string{"t"},
						Value:       "",
						Usage:       "Enter what you want to find",
						Destination: &interaction.CurrentRequest,
					},
				},
				Category: "Storage",
				Action:   find,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func keep(ctx context.Context, c *cli.Command) error {
	fmt.Print("hello:" + interaction.CurrentRequest)
	return nil
}

func ponder(ctx context.Context, c *cli.Command) error {
	if interaction.CurrentRequest == "" {
		return fmt.Errorf("topic is required")
	}

	content, err := QueryLLM(interaction.CurrentRequest)
	if err != nil {
		return fmt.Errorf("failed to query LLM: %w", err)
	}

	// Format and display the response
	formattedResponse := text.WrapSoft(content, 150)
	coloredResponse := text.Colors{text.FgBlue}.Sprint(formattedResponse)
	fmt.Println(coloredResponse)

	return nil
}
