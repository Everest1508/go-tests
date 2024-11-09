package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/everest1508/go-tests-setup/auto/utils"
	"github.com/playwright-community/playwright-go"
)

func loadConfig(filePath string) (*utils.TestConfig, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filePath, err)
	}
	var config utils.TestConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("could not parse JSON: %v", err)
	}
	return &config, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path_to_config.json>")
	}
	configFile := os.Args[1]
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(*config.Browser)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	utils.ExecuteTestSteps(page, config.Tests)

	fmt.Println("All test steps completed successfully.")
}
