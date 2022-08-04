package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/urfave/cli/v2"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

const filename = ".openai.yml"

type config struct {
	APIKey string `yaml:"api-key"`
}

func loadConfigFile() (config, error) {
	oldFile, err := readFile()
	if err != nil {
		return config{}, fmt.Errorf("unable to read file: %w", err)
	}

	var oldConfig config

	err = yaml.Unmarshal(oldFile, &oldConfig)
	if err != nil {
		return config{}, fmt.Errorf("unable to unmarshal file: %w", err)
	}

	return oldConfig, nil
}

func initClient(c *cli.Context) error {
	fmt.Print("Enter OPEN API_KEY: ")

	apiKey, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read api_key: %w", err)
	}

	oldConfig, err := loadConfigFile()
	if err != nil {
		oldConfig = config{}
	}

	oldConfig.APIKey = string(apiKey)

	data, err := yaml.Marshal(&oldConfig)
	if err != nil {
		return fmt.Errorf("unable to marshal yaml: %w", err)
	}

	if err := writeFile(data); err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %w", err)
	}

	return filepath.Join(home, filename), nil
}

func readFile() ([]byte, error) {
	fp, err := getFilePath()
	if err != nil {
		return nil, fmt.Errorf("unable to get file path: %w", err)
	}

	oldFile, err := os.ReadFile(fp)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("unable to read file: %w", err)
		}
	}

	return oldFile, nil
}

func writeFile(data []byte) error {
	fp, err := getFilePath()
	if err != nil {
		return fmt.Errorf("unable to get file path: %w", err)
	}

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}
