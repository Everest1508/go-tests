package e2e

import (
	"fmt"

	"github.com/everest1508/go-tests-setup/scripts"
	"github.com/playwright-community/playwright-go"

	pages "github.com/everest1508/go-tests-setup/internal/pages/auth"

	"testing"
)

func TestRegister(t *testing.T) {
	ts := scripts.Test{Headless: false, SlowMo: playwright.Float(1000)}
	pw, browser, logger, _ := ts.TestSetup()
	defer ts.TestTeardown(pw, browser)

	page, err := browser.NewPage()
	if err != nil {
		logger.Error(fmt.Sprintf("could not create page: %v", err))
	}

	lf := pages.RegisterForm{
		Email:    "ritesh@scalet.io",
		Password: "Scalent@123",
	}

	rp := pages.NewRegistrationPage(page)
	rp.Navigate()
	rp.Register(lf)
}

func TestLogin(t *testing.T) {
	ts := scripts.Test{Headless: false, SlowMo: playwright.Float(1000)}
	pw, browser, logger, _ := ts.TestSetup()
	defer ts.TestTeardown(pw, browser)

	page, err := browser.NewPage()
	if err != nil {
		logger.Error(fmt.Sprintf("could not create page: %v", err))
	}

	lf := pages.LoginForm{
		Email:    "ritesh@scalet.io",
		Password: "Scalent@123",
	}

	loginPage := pages.NewLoginPage(page)
	loginPage.Navigate()
	loginPage.Login(lf)
}
