package stepfunctions

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"playwright-golang-practice/helpers"
	//"playwright-golang-practice/stepdefinitions/models"
)

//type ExtendedEntity struct {
//	x *models.Entity
//}

type Entity struct {
	Pw       *playwright.Playwright
	Browser  playwright.Browser
	Page     playwright.Page
	FullName string
	Cases    helpers.Case
}

func (e *Entity) IOpenTheWebsite(params string) error {
	fmt.Println("starting")
	pw, err := playwright.Run()
	if err != nil {
		return fmt.Errorf("playwright: %w", err)
	}
	err = helpers.LogPanic(err)
	if err != nil {
		return err
	}
	e.Pw = pw
	fmt.Println("finisihed")

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	err = helpers.LogPanic(err)
	if err != nil {
		return err
	}
	e.Browser = browser

	page, err := browser.NewPage()
	err = helpers.LogPanic(err)
	if err != nil {
		return err
	}

	err = page.SetViewportSize(1920, 1440)
	err = helpers.LogPanic(err)
	if err != nil {
		return err
	}

	e.Page = page

	_, err = page.Goto(params)
	err = helpers.LogPanic(err)
	if err != nil {
		return err
	}
	return nil
}
