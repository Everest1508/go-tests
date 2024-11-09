package utils

import (
	"fmt"

	"github.com/everest1508/go-tests-setup/auto/constants"
	"github.com/playwright-community/playwright-go"
)

func ExecuteTestSteps(page playwright.Page, steps []TestStep) error {
	var err error
	for _, step := range steps {
		switch step.SelectorType {
		case constants.Goto:
			err = handleGotoAction(page, step)
		case constants.Role:
			err = handleRoleAction(page, step)
		case constants.Default:
			err = handleDefaultAction(page, step)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func handleRoleAction(page playwright.Page, step TestStep) error {
	switch step.Action {
	case "click":
		return page.GetByRole(playwright.AriaRole(step.Role.Type), *step.Role.Options).Click()
	case "fill":
		if step.Value == nil {
			return fmt.Errorf("no value provided for fill action")
		}
		return page.GetByRole(playwright.AriaRole(step.Role.Type), *step.Role.Options).Fill(*step.Value)
	default:
		return fmt.Errorf("unknown action: %s", step.Action)
	}
}

func handleGotoAction(page playwright.Page, step TestStep) error {
	fmt.Println(*step.Value)
	_, err := page.Goto(*step.Value)
	return err
}

func handleDefaultAction(page playwright.Page, step TestStep) error {
	switch step.Action {
	case "click":
		return page.Locator(step.Selector).Click()
	case "fill":
		if step.Value == nil {
			return fmt.Errorf("no value provided for fill action")
		}
		return page.Locator(step.Selector).Fill(*step.Value)
	default:
		return fmt.Errorf("unknown action: %s", step.Action)
	}
}
