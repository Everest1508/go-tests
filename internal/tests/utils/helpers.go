package utils

import (
	"time"

	"github.com/playwright-community/playwright-go"
)

func WaitForSelector(page playwright.Page, selector string, timeout int) (playwright.ElementHandle, error) {
	return page.WaitForSelector(selector, playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(float64(timeout)),
	})
}

func ClickElement(page playwright.Page, selector string, timeout int) error {
	if _, err := WaitForSelector(page, selector, timeout); err != nil {
		return err
	}
	return page.Click(selector)
}

func FillInput(page playwright.Page, selector string, value string, timeout int) error {
	if _, err := WaitForSelector(page, selector, timeout); err != nil {
		return err
	}
	return page.Fill(selector, value)
}

func HandleAlert(page playwright.Page, timeout int) (string, error) {
	var alertMessage string
	page.On("dialog", func(dialog playwright.Dialog) {
		alertMessage = dialog.Message()
		dialog.Accept()
	})

	time.Sleep(time.Duration(timeout) * time.Millisecond)
	return alertMessage, nil
}

func NavigateTo(page playwright.Page, url string) (playwright.Response, error) {
	return page.Goto(url)
}
