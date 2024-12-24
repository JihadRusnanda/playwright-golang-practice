package models

import (
	"github.com/playwright-community/playwright-go"
	"playwright-golang-practice/helpers"
)

type Entity struct {
	Pw       *playwright.Playwright
	Browser  playwright.Browser
	Page     playwright.Page
	FullName string
	Cases    helpers.Case
}
