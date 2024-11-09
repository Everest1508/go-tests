package scripts

import (
	"fmt"
	"log"

	"github.com/everest1508/go-tests-setup/internal/utils"
	"github.com/playwright-community/playwright-go"
)

type Test struct {
	Headless bool
	SlowMo   *float64
}

func (t *Test) TestSetup() (*playwright.Playwright, playwright.Browser, *utils.Logger, error) {
	logger := utils.NewLogger()

	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not start Playwright: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(t.Headless),
		SlowMo:   t.SlowMo,
	})
	if err != nil {
		pw.Stop()
		return nil, nil, nil, fmt.Errorf("could not launch browser: %w", err)
	}

	return pw, browser, logger, nil
}

func (t *Test) TestTeardown(pw *playwright.Playwright, browser playwright.Browser) {
	if browser != nil {
		if err := browser.Close(); err != nil {
			log.Printf("Error closing browser: %v", err)
		}
	}
	if pw != nil {
		pw.Stop()
	}
}
