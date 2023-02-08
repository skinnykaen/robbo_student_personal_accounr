package models

type DataTest struct {
	Name          string
	Token         string
	Variables     map[string]interface{}
	ExpectedError error
}
