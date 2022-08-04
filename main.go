package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mbenaiss/openai/cli/openai"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize openai",
				Action:  initClient,
			},
		},
		Name:  "OpenAI",
		Usage: "OpenAI technology is a cutting edge technology that allows for artificial intelligence to be used in a variety of ways.",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "prompt", Value: "", Usage: "prompt", Aliases: []string{"p"}},
			&cli.Float64Flag{Name: "temperature", Value: 0.5, Usage: "temperature", Aliases: []string{"t"}},
			&cli.IntFlag{Name: "token", Value: 100, Usage: "token", Aliases: []string{"mt"}},
			&cli.Float64Flag{Name: "frequency-penalty", Value: 0.0, Usage: "frequency penalty", Aliases: []string{"fp"}},
			&cli.Float64Flag{Name: "presence-penalty", Value: 0.0, Usage: "presence penalty", Aliases: []string{"pp"}},
		},
		Action: func(c *cli.Context) error {
			config, err := loadConfigFile()
			if err != nil {
				return err
			}

			openaiClient := openai.New(config.APIKey)

			payload := openai.Payload{
				Model:            "text-davinci-002",
				Prompt:           c.String("prompt"),
				Temperature:      c.Float64("temperature"),
				MaxTokens:        c.Int("token"),
				TopP:             1,
				FrequencyPenalty: c.Float64("frequency-penalty"),
				PresencePenalty:  c.Float64("presence-penalty"),
			}

			text, err := openaiClient.Request(payload)
			if err != nil {
				return fmt.Errorf("unable to request: %w", err)
			}

			fmt.Println(text)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
