package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mbenaiss/openai/cli/openai"
)

var model = map[string]string{
	"curie":   "text-curie-001",
	"davinci": "text-davinci-002",
	"codex":   "code-davinci-002",
}

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
			&cli.StringFlag{Name: "model", Value: "davinci", Usage: "model", Aliases: []string{"m"}},
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

			model := model[c.String("model")]
			if model == "" {
				model = "text-davinci-002"
			}

			payload := openai.Payload{
				Model:            model,
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
