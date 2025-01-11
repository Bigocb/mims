package main

import (
	"fmt"
	"jervis/external"
	"jervis/tui"
	"log"
	"os"

	"github.com/urfave/cli"
)

func buildQuery(question string) string {

	var resp string
	resp = external.QueryLLM(question)
	return resp
}

func main() {
	app := cli.NewApp()
	app.Name = ""
	app.Usage = ""

	// We'll be using the same flag for all our commands
	// so we'll define it up here
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "topic",
			Value: "Please enter a topic",
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
				content := ns
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
