// test_page.go.tpl
package main

import (
	"github.com/playwright-community/playwright-go"
)

type LoginForm struct {
	Username string
	Password string
}

type LoginPage struct {
	page playwright.Page
}

func NewLoginPage(page playwright.Page) *LoginPage {
	return &LoginPage{page: page}
}

func (p *LoginPage) RunTest(form LoginForm) {

	p.page.Goto("https://practicetestautomation.com/practice-test-login/")

	p.page.Fill("input[name=&#39;username&#39;]", form.Username)

	p.page.Fill("input[name=&#39;password&#39;]", form.Password)

	p.page.GetByRole("button", playwright.PageGetByRoleOptions{
		Name:  "",
		Exact: playwright.Bool(true),
	}).Click()

}
