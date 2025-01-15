package main

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v3"
	"jervis/llm"
	"log"
	"os"
)

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

func main() {

	var topic string

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "topic",
				Value:       "list of companies like solo.io",
				Usage:       "Please enter a topic",
				Destination: &topic,
			},
		},
		Name:  "research",
		Usage: "Get response from AI",
		Action: func(ctx context.Context, c *cli.Command) error {
			if topic == "" {
				return fmt.Errorf("topic is required")
			}

			content, err := QueryLLM(topic)
			if err != nil {
				return fmt.Errorf("failed to query LLM: %w", err)
			}

			// Format and display the response
			formattedResponse := text.WrapSoft(content, 150)
			coloredResponse := text.Colors{text.FgBlue}.Sprint(formattedResponse)
			fmt.Println(coloredResponse)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
