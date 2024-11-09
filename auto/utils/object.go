package utils

import (
	"github.com/everest1508/go-tests-setup/auto/constants"
	"github.com/playwright-community/playwright-go"
)

type Role struct {
	Type    string                           `json:"type"`
	Options *playwright.PageGetByRoleOptions `json:"options"`
}

type TestStep struct {
	Name         *string                `json:"name"`
	Action       constants.Action       `json:"action"`
	SelectorType constants.SelectorType `json:"selectorType,omitempty"`
	Selector     string                 `json:"selector,omitempty"`
	Value        *string                `json:"value,omitempty"`
	AssertType   string                 `json:"assertType,omitempty"`
	Role         Role                   `json:"role,omitempty"`
}

type TestConfig struct {
	TestName string                               `json:"testName"`
	Browser  *playwright.BrowserTypeLaunchOptions `json:"browser"`
	Tests    []TestStep                           `json:"tests"`
}
