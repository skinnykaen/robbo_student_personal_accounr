package integration_tests

import (
	"context"
	"github.com/skinnykaen/robbo_student_personal_account.git/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app, cleanerContainer := app.TestApp()
	app.Start(context.Background())
	code := m.Run()
	app.Stop(context.Background())
	cleanerContainer()
	os.Exit(code)
}
