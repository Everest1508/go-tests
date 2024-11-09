// test_page.go.tpl
package pages

import (
	"github.com/playwright-community/playwright-go"
)

type {{.TestName}}Form struct {
	{{- range .Steps }}
		{{- if eq .Action "fill" }}
		{{.Name}} string
		{{- end }}
	{{- end }}
}

type {{.TestName}}Page struct {
	page playwright.Page
}

func New{{.TestName}}Page(page playwright.Page) *{{.TestName}}Page {
	return &{{.TestName}}Page{page: page}
}

func (p *{{.TestName}}Page) RunTest(form {{.TestName}}Form) {
	{{range .Steps}}
	{{- if eq .Action "fill" }}
	p.page.Fill("{{.Selector}}", form.{{.Name}})
	{{- else if eq .Action "click" }}
		{{- if eq .SelectorType "role" }}
		p.page.GetByRole("{{.Role.Type}}", playwright.PageGetByRoleOptions{
			Name:  "{{.Role.Options.Name}}",
			Exact: playwright.Bool({{.Role.Options.Exact}}),
		}).Click()
		{{- else }}
		p.page.Click("{{.Selector}}")
		{{- end }}
	{{- else if eq .Action "goto" }}
		p.page.Goto("{{.Value}}")
	{{- else if eq .Action "assert" }}
		{{- if eq .AssertType "text" }}
		
		{{- else if eq .AssertType "url" }}
		
		{{- end }}
	{{- end }}
	{{end}}
}
