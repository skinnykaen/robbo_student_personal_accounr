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

func TestGetRobboUnitById(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	testTable := factory.TestTableGetRobboUnitById()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var q struct {
				GetRobboUnitById struct {
					models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
					models.Error         `graphql:"... on Error"`
				} `graphql:"GetRobboUnitById(id: $RobboUnitId)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, q.GetRobboUnitById.ID, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).ID)
				assert.Equal(t, q.GetRobboUnitById.Name, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).Name)
				assert.Equal(t, q.GetRobboUnitById.City, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).City)
			}
		})
	}
}

func TestGetAllRobboUnits(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	testTable := factory.TestTableGetAllRobboUnits()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var q struct {
				GetAllRobboUnits struct {
					models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
					models.Error             `graphql:"... on Error"`
				} `graphql:"GetAllRobboUnits(page: $Page, pageSize: $PageSize)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, len(q.GetAllRobboUnits.RobboUnits), len(testCase.Body["RobboUnits"].([]*models.RobboUnitHTTP)))
				assert.Equal(t, q.GetAllRobboUnits.CountRows, testCase.Body["CountRows"])
			}
		})
	}
}

func TestGetRobboUnitsByAccessToken(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	testTable := factory.TestTableGetRobboUnitsByAccessToken()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var q struct {
				GetRobboUnitsByAccessToken struct {
					models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
					models.Error             `graphql:"... on Error"`
				} `graphql:"GetRobboUnitsByAccessToken(page: $Page, pageSize: $PageSize)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, len(q.GetRobboUnitsByAccessToken.RobboUnits), len(testCase.Body["RobboUnits"].([]*models.RobboUnitHTTP)))
				assert.Equal(t, q.GetRobboUnitsByAccessToken.CountRows, testCase.Body["CountRows"])
			}
		})
	}
}

func TestGetRobboUnitsByUnitAdminId(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	testTable := factory.TestTableGetRobboUnitsByUnitAdminId()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var q struct {
				GetRobboUnitsByUnitAdminId struct {
					models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
					models.Error             `graphql:"... on Error"`
				} `graphql:"GetRobboUnitsByUnitAdminId(unitAdminId: $UnitAdminId)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, len(q.GetRobboUnitsByUnitAdminId.RobboUnits), len(testCase.Body["RobboUnits"].([]*models.RobboUnitHTTP)))
			}
		})
	}
}

func TestSearchRobboUnitsByName(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//httpClient := &http.Client{Transport: tr}
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	testTable := factory.TestTableSearchRobboUnitsByName()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var q struct {
				SearchRobboUnitsByName struct {
					models.RobboUnitHTTPList `graphql:"... on RobboUnitHttpList"`
					models.Error             `graphql:"... on Error"`
				} `graphql:"SearchRobboUnitsByName(name: $Name, page: $Page, pageSize: $PageSize)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Query(context.Background(), &q, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, len(q.SearchRobboUnitsByName.RobboUnits), len(testCase.Body["RobboUnits"].([]*models.RobboUnitHTTP)))
				assert.Equal(t, q.SearchRobboUnitsByName.CountRows, testCase.Body["CountRows"])
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

	testTable := factory.TestTableUpdateRobboUnit()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {

			var m struct {
				UpdateRobboUnit struct {
					models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
					models.Error         `graphql:"... on Error"`
				} `graphql:"UpdateRobboUnit(input: $UpdateRobboUnit)"`
			}

			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Mutate(context.Background(), &m, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, m.UpdateRobboUnit.ID, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).ID)
				assert.Equal(t, m.UpdateRobboUnit.Name, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).Name)
				assert.Equal(t, m.UpdateRobboUnit.City, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).City)
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

	testTable := factory.TestTableDeleteRobboUnit()

	for _, testCase := range testTable {

		var m struct {
			DeleteRobboUnit struct {
				models.DeletedRobboUnit `graphql:"... on DeletedRobboUnit"`
			} `graphql:"DeleteRobboUnit(robboUnitId: $RobboUnitId)"`
		}

		t.Run(testCase.Name, func(t *testing.T) {
			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Mutate(context.Background(), &m, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, m.DeleteRobboUnit.DeletedRobboUnit.RobboUnitID, testCase.Body["DeletedRobboUnit"].(models.DeletedRobboUnit).RobboUnitID)
			}
		})
	}
}

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

	testTable := factory.TestTableCreateRobboUnit()

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			expect := testCase.ExpectedError
			gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
				request.Header.Add("Authorization", testCase.Token)
			})
			correct := gqlClientWithRequestModifier.Mutate(context.Background(), &m, testCase.Variables)
			if correct != nil {
				assert.Equal(t, expect.Error(), correct.Error())
			} else {
				assert.Equal(t, m.CreateRobboUnit.RobboUnitHTTP.ID, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).ID)
				assert.Equal(t, m.CreateRobboUnit.RobboUnitHTTP.Name, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).Name)
				assert.Equal(t, m.CreateRobboUnit.RobboUnitHTTP.City, testCase.Body["RobboUnit"].(models.RobboUnitHTTP).City)
			}
		})
	}
}
