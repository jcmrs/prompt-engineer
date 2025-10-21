package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcmrs/prompt-engineer/internal/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pea",
		Usage: "Prompt Engineering Agent",
		Commands: []*cli.Command{
			{
				Name:  "check-auth",
				Usage: "Check Gemini CLI authentication",
				Action: func(c *cli.Context) error {
					fmt.Println("Not implemented")
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Start the PEA server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "db",
						Value: "",
						Usage: "Path to the database file",
					},
					&cli.IntFlag{
						Name:  "port",
						Value: 8080,
						Usage: "Port to listen on",
					},
				},
				Action: func(c *cli.Context) error {
					s := server.NewServer()
					s.Addr = fmt.Sprintf(":%d", c.Int("port"))
					return s.ListenAndServe()
				},
			},
			{
				Name:  "run",
				Usage: "Run a prompt template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "template",
						Required: true,
						Usage:    "ID of the prompt template to run",
					},
					&cli.BoolFlag{
						Name:  "no-store",
						Usage: "Do not persist the run",
					},
					&cli.BoolFlag{
						Name:  "allow-local-tests",
						Usage: "Allow running local tests",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Not implemented")
					return nil
				},
			},
			{
				Name:  "export",
				Usage: "Export a conversation",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "conversation",
						Required: true,
						Usage:    "ID of the conversation to export",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: "markdown",
						Usage: "Export format (markdown)",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Not implemented")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
