package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"jervis/external"
	"jervis/tui"
	"log"
	"os"

	"github.com/urfave/cli"
)

func buildQuery(question string) openai.ChatCompletionResponse {
	message := []openai.ChatCompletionMessage{
		openai.ChatCompletionMessage{
			Role:    "system",
			Content: "You are a helpful AI, named mim",
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: message,
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})

	var resp openai.ChatCompletionResponse
	resp = external.QueryOpenAi(req)
	return resp
}

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	// We'll be using the same flag for all our commands
	// so we'll define it up here
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Value: "tutorialedge.net",
		},
	}
	var err error
	// we create our commands
	app.Commands = []cli.Command{
		{
			Name:  "research",
			Usage: "Get response from AI",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				test := c.String("topic")
				fmt.Println(test)
				ns := buildQuery(c.String("topic"))
				if err != nil {
					return err
				}
				content := ns.Choices[0].Message.Content
				fmt.Println(content)
				return nil
			},
		},
		{
			Name:  "chat",
			Usage: "Open Interactive Chat",
			Action: func(c *cli.Context) error {
				// run main function from tui package
				tui.Tui()
				return nil
			},
		},
	}

	// start our application
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
