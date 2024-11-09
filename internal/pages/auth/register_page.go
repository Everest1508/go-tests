package pages

import (
	"github.com/playwright-community/playwright-go"
)

type RegisterForm struct {
	Email    string
	Password string
}

type RegisterPage struct {
	page playwright.Page
}

func NewRegistrationPage(page playwright.Page) *RegisterPage {
	return &RegisterPage{page: page}
}

func (p *RegisterPage) Navigate() {
	p.page.Goto("https://bulkcalendar.rubaru.shop/auth/register")
}

func (p *RegisterPage) Register(f RegisterForm) {
	p.page.GetByPlaceholder("Email", playwright.PageGetByPlaceholderOptions{
		Exact: playwright.Bool(true),
	}).Fill(f.Email)
	p.page.GetByPlaceholder("New Password", playwright.PageGetByPlaceholderOptions{
		Exact: playwright.Bool(true),
	}).Fill(f.Password)
	p.page.GetByRole("input", playwright.PageGetByRoleOptions{Checked: playwright.Bool(false)}).Click()
	p.page.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Create account", Exact: playwright.Bool(true)})
}
