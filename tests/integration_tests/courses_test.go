package integration_tests

import (
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/tests/factory"
	"net/http"
	"testing"
)

func TestCreateAccessCourseRelationRobboGroup(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	var m struct {
		CreateAccessCourseRelationRobboGroup struct {
			models.CourseRelationHTTP `graphql:"... on CourseRelationHttp"`
			models.Error              `graphql:"... on Error"`
		} `graphql:"CreateAccessCourseRelationRobboGroup(input: $NewCourseAccessRelationRobboGroup)"`
	}

	testData := factory.DataCreateCourseAccessRelationRobboGroup()

	for _, testCase := range testData {
		t.Run(testCase.Name, func(t *testing.T) {
			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Mutate(context.Background(), &m, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, expect, correct)
			}
		})
	}
}
