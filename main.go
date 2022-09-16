package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/mbenaiss/openai/cli/openai"
)

var model = map[string]string{
	"curie":   "text-curie-001",
	"davinci": "text-davinci-002",
	"codex":   "code-davinci-002",
}

var (
	modelFlag       = cli.StringFlag{Name: "model", Value: "davinci", Usage: "model", Aliases: []string{"m"}}
	temperatureFlag = cli.Float64Flag{Name: "temperature", Value: 0.5, Usage: "temperature", Aliases: []string{"t"}}
	tokenFlag       = cli.IntFlag{Name: "token", Value: 100, Usage: "token", Aliases: []string{"mt"}}
	fpFlag          = cli.Float64Flag{Name: "frequency-penalty", Value: 0.0, Usage: "frequency penalty", Aliases: []string{"fp"}}
	ppFlag          = cli.Float64Flag{Name: "presence-penalty", Value: 0.0, Usage: "presence penalty", Aliases: []string{"pp"}}
	stopFlag        = cli.StringSliceFlag{Name: "stop", Value: cli.NewStringSlice(), Usage: "stop", Aliases: []string{"s"}}
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
			{
				Name:    "codex",
				Aliases: []string{"c"},
				Usage:   "Generate code",
				Flags:   []cli.Flag{&temperatureFlag, &tokenFlag, &fpFlag, &ppFlag, &stopFlag},
				Action: func(c *cli.Context) error {
					stop := c.StringSlice("stop")
					if len(stop) == 0 {
						stop = []string{"OUTPUT", "</code>"}
					}

					promptArgs := c.Args().Get(0)

					prompt := fmt.Sprintf("%s \nCODE: \n", promptArgs)
					payload := openai.Payload{
						Prompt:           prompt,
						Model:            model["codex"],
						Temperature:      c.Float64("temperature"),
						TopP:             1,
						MaxTokens:        c.Int("token"),
						FrequencyPenalty: c.Float64("frequency-penalty"),
						PresencePenalty:  c.Float64("presence-penalty"),
						Stop:             stop,
					}

					text, err := callOpenAI(payload)
					if err != nil {
						return err
					}

					text = strings.TrimLeft(text, "<code>")

					fmt.Println(text)

					return nil
				},
			},
		},
		Name:      "OpenAI",
		Usage:     "OpenAI technology is a cutting edge technology that allows for artificial intelligence to be used in a variety of ways.",
		Flags:     []cli.Flag{&modelFlag, &temperatureFlag, &tokenFlag, &fpFlag, &ppFlag, &stopFlag},
		ArgsUsage: "[prompt]",
		Action: func(c *cli.Context) error {
			model := model[c.String("model")]
			if model == "" {
				model = "text-davinci-002"
			}

			stop := c.StringSlice("stop")
			if len(stop) == 0 {
				stop = nil
			}

			payload := openai.Payload{
				Model:            model,
				Prompt:           c.Args().Get(0),
				Temperature:      c.Float64("temperature"),
				MaxTokens:        c.Int("token"),
				TopP:             1,
				FrequencyPenalty: c.Float64("frequency-penalty"),
				PresencePenalty:  c.Float64("presence-penalty"),
				Stop:             stop,
			}

			text, err := callOpenAI(payload)
			if err != nil {
				return err
			}

			fmt.Println(text)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func callOpenAI(payload openai.Payload) (string, error) {
	config, err := loadConfigFile()
	if err != nil {
		return "", err
	}

	openaiClient := openai.New(config.APIKey)

	text, err := openaiClient.Request(payload)
	if err != nil {
		return "", fmt.Errorf("unable to request: %w", err)
	}

	return text, nil
}
