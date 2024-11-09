package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type TestConfig struct {
	TestName string  `json:"testName"`
	Browser  Browser `json:"browser"`
	Tests    []Test  `json:"tests"`
}

type Browser struct {
	Headless bool `json:"headless"`
	SlowMo   int  `json:"slowMo"`
}

type Test struct {
	Action       string `json:"action"`
	SelectorType string `json:"selectorType"`
	Selector     string `json:"selector"`
	Value        string `json:"value,omitempty"`
	Role         *Role  `json:"role,omitempty"`
}

type Role struct {
	Type    string      `json:"type"`
	Options RoleOptions `json:"options"`
}

type RoleOptions struct {
	Name  string `json:"name"`
	Exact bool   `json:"exact"`
}

func main() {
	// Set the URL you want to scrape
	url := "https://practicetestautomation.com/practice-test-login/"

	// Create a custom HTTP client with an increased timeout
	client := &http.Client{
		Timeout: 30 * time.Second, // Increase timeout to 30 seconds
	}

	// Fetch the page
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML from the response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing the HTML:", err)
		return
	}

	// Initialize TestConfig
	config := TestConfig{
		TestName: "Login",
		Browser: Browser{
			Headless: false,
			SlowMo:   1000,
		},
		Tests: []Test{},
	}

	// Loop through all the form elements and extract their details
	// doc.Find("form").Each(func(i int, form *goquery.Selection) {
	// 	// Extract and print the input fields inside the form
	doc.Find("input, button, select").Each(func(j int, el *goquery.Selection) {
		tag := el.Get(0).Data
		name, _ := el.Attr("name")
		// id, _ := el.Attr("id")
		value, _ := el.Attr("value")
		// placeholder, _ := el.Attr("placeholder")

		// Handle inputs
		if tag == "input" {
			if el.AttrOr("type", "") == "submit" || el.AttrOr("type", "") == "button" {
				// Handle button-like elements
				config.Tests = append(config.Tests, Test{
					Action:       "click",
					SelectorType: "default",
					Selector:     fmt.Sprintf("button[name='%s']", name),
				})
			} else {
				// Handle other input fields
				config.Tests = append(config.Tests, Test{
					Action:       "fill",
					SelectorType: "default",
					Selector:     fmt.Sprintf("input[name='%s']", name),
					Value:        value,
				})
			}
		} else if tag == "select" {
			// Handle select dropdowns
			config.Tests = append(config.Tests, Test{
				Action:       "fill",
				SelectorType: "default",
				Selector:     fmt.Sprintf("select[name='%s']", name),
			})
		} else if tag == "button" {
			// Handle buttons
			config.Tests = append(config.Tests, Test{
				Action:       "click",
				SelectorType: "role",
				Selector:     fmt.Sprintf("button[name='%s']", name),
				Role: &Role{
					Type: "button",
					Options: RoleOptions{
						Name:  name,
						Exact: true,
					},
				},
			})
		}
	})
	// })

	// Output the result as JSON
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Save to a file
	err = os.WriteFile("test_config.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("Test configuration saved to test_config.json")
}
