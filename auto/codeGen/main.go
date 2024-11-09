package main

import "log"

func main() {
	config, err := loadTestConfig("../login_test.json")
	if err != nil {
		log.Fatalf("Error loading test config: %v", err)
	}

	templateData := prepareTemplateData(config)

	if err := generateCode(templateData); err != nil {
		log.Fatalf("Error generating code: %v", err)
	}
}
