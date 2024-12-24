package stepdefinitions

import (
	"context"
	"github.com/cucumber/godog"
	"playwright-golang-practice/helpers"
	"playwright-golang-practice/stepfunctions"
)

func InitializeScenario(sc *godog.ScenarioContext) {

	step := &stepfunctions.Entity{}

	sc.Step(`^I open website "([^"]*)"$`, step.IOpenTheWebsite)
	//sc.Step(`I fill form in following information:`, step.IFillForm)
	//sc.Step(`I click submit button`, step.IClickSubmit)
	//sc.Step(`Verify result information "([^"]*)"`, step.VerifyResult)

	sc.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		if err := step.Page.Close(); err != nil {
			err := helpers.LogPanic(err)
			if err != nil {
				return nil, err
			}
		}
		if err := step.Browser.Close(); err != nil {
			err := helpers.LogPanic(err)
			if err != nil {
				return nil, err
			}
		}
		if err := step.Pw.Stop(); err != nil {
			err := helpers.LogPanic(err)
			if err != nil {
				return nil, err
			}
		}
		return ctx, nil
	})
}
