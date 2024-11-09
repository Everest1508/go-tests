package pages

import (
	"github.com/playwright-community/playwright-go"
)

type LoginForm struct {
	Email    string
	Password string
}

type LoginPage struct {
	page playwright.Page
}

func NewLoginPage(page playwright.Page) *LoginPage {
	return &LoginPage{page: page}
}

func (lp *LoginPage) Navigate() {
	lp.page.Goto("https://bulkcalendar.rubaru.shop/auth/login")
}

func (lp *LoginPage) Login(lf LoginForm) {
	lp.page.Fill("input[type='email']", lf.Email)
	lp.page.Fill("input[type='password']", lf.Password)
	lp.page.GetByRole("button", playwright.PageGetByRoleOptions{
		Name:  "Sign in",
		Exact: playwright.Bool(true),
	}).Click()
}
