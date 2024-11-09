package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/everest1508/go-tests-setup/auto/utils"
	"github.com/playwright-community/playwright-go"
)

type TemplateData struct {
	TestName string
	Browser  *playwright.BrowserTypeLaunchOptions
	Steps    []utils.TestStep
}

func loadTestConfig(filename string) (*utils.TestConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var testConfig utils.TestConfig
	if err := json.Unmarshal(file, &testConfig); err != nil {
		return nil, err
	}
	return &testConfig, nil
}

func prepareTemplateData(config *utils.TestConfig) *TemplateData {
	return &TemplateData{
		TestName: config.TestName,
		Browser:  config.Browser,
		Steps:    config.Tests,
	}
}

func generateCode(templateData *TemplateData) error {
	tmpl, err := template.ParseFiles("../templates/test_page.go.tpl")
	if err != nil {
		return err
	}

	file, err := os.Create("generated_page.go")
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, templateData)
}
