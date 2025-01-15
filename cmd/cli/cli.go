package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"jervis/llm"
	"log"
	"os"
)

/*
Current this is a very simple cli that takes in a question and returns an LLM response.
*/

// buildQuery
func buildQuery(question string) string {
	request := llm.Request{
		Query: question,
	}
	var resp llm.Response
	resp = llm.Query(request)
	return resp.Details
}

func main() {

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "topic",
			Value: "Please enter a topic",
		},
	}

	var err error
	// we create our commands
	cmd := &cli.Command{
		Name:  "research",
		Usage: "Get response from AI",
		Flags: myFlags,
		// the action, or code that will be executed when
		// we execute our `ns` command
		Action: func(ctx context.Context, c *cli.Command) error {
			// a simple lookup function
			test := c.String("topic")
			fmt.Println(test)
			ns := buildQuery(c.String("topic"))
			if err != nil {
				return err
			}
			content := ns
			fmt.Println(content)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
