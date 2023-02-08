package integration_tests

import (
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/tests/factory"
	"log"
	"net/http"
	"testing"
)

func TestCreateRobboUnit(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	var m struct {
		CreateRobboUnit struct {
			models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
			models.Error         `graphql:"... on Error"`
		} `graphql:"CreateRobboUnit(input: $NewRobboUnit)"`
	}

	testData := factory.DataCreateRobboUnit()

	for _, testCase := range testData {
		t.Run(testCase.Name, func(t *testing.T) {
			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Mutate(context.Background(), &m, testCase.Variables)
			if correct != nil {
				log.Println(correct)
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, expect, correct)
			}
		})
	}
}

func TestUpdateRobboUnit(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	var m struct {
		UpdateRobboUnit struct {
			models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
			models.Error         `graphql:"... on Error"`
		} `graphql:"UpdateRobboUnit(input: $UpdateRobboUnit)"`
	}

	testData := factory.DataUpdateRobboUnit()

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

func TestGetRobboUnitById(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	var q struct {
		GetRobboUnitById struct {
			models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
			models.Error         `graphql:"... on Error"`
		} `graphql:"GetRobboUnitById(id: $RobboUnitId)"`
	}

	testData := factory.DataGetRobboUnitById()

	for _, testCase := range testData {
		t.Run(testCase.Name, func(t *testing.T) {
			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, expect, correct)
			}
		})
	}
}

func TestDeleteRobboUnit(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	var m struct {
		DeleteRobboUnit struct {
			models.DeletedRobboUnit `graphql:"... on DeletedRobboUnit"`
		} `graphql:"DeleteRobboUnit(robboUnitId: $RobboUnitId)"`
	}

	testData := factory.DataDeleteRobboUnit()

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

//
//func TestGetAllRobboUnits(t *testing.T) {
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	httpClient := &http.Client{Transport: tr}
//httpClient := &http.Client{}
//gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)
//
//	var m struct {
//		GetAllRobboUnits struct {
//			models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
//			models.Error         `graphql:"... on Error"`
//		} `graphql:"GetAllRobboUnits"`
//	}
//}
//
//func TestGetRobboUnitsByUnitAdminId(t *testing.T) {
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	httpClient := &http.Client{Transport: tr}
//httpClient := &http.Client{}
//gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)
//
//	var m struct {
//		GetRobboUnitsByUnitAdminId struct {
//			models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
//			models.Error         `graphql:"... on Error"`
//		} `graphql:"GetRobboUnitsByUnitAdminId(unitAdminId: $unitAdminId)"`
//	}
//}
//
//func TestGetRobboUnitsByAccessToken(t *testing.T) {
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	httpClient := &http.Client{Transport: tr}
//httpClient := &http.Client{}
//gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)
//
//	var m struct {
//		GetRobboUnitsByAccessToken struct {
//			models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
//			models.Error         `graphql:"... on Error"`
//		} `graphql:"GetRobboUnitsByAccessToken"`
//	}
//}
