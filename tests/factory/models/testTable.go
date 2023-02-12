package models

type TestTable struct {
	Name          string
	Token         string
	Body          map[string]interface{}
	Variables     map[string]interface{}
	ExpectedError error
}
