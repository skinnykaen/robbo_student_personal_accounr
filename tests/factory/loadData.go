package factory

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func LoadData() {
	LoadDataForTestUnitAdmin()
	LoadDataForTestRobboUnit()
}

func LoadDataForTestUnitAdmin() {
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	unitAdmin := models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      "unitadmin@test.ru",
			Password:   "test",
			Role:       4,
			Nickname:   "test",
			Firstname:  "test",
			Lastname:   "test",
			Middlename: "test",
		},
	}

	gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
		request.Header.Add("Authorization", "Bearer "+viper.GetString("auth.tokens.super_admin"))
	})

	var m struct {
		CreateUnitAdmin struct {
			models.UnitAdminHTTP `graphql:"... on UnitAdminHttp"`
			models.Error         `graphql:"... on Error"`
		} `graphql:"CreateUnitAdmin(input: $NewUnitAdmin)"`
	}

	variables := map[string]interface{}{
		"NewUnitAdmin": models.NewUnitAdmin{
			Email:      unitAdmin.UserHTTP.Email,
			Password:   unitAdmin.UserHTTP.Email,
			Nickname:   unitAdmin.UserHTTP.Nickname,
			Firstname:  unitAdmin.UserHTTP.Firstname,
			Lastname:   unitAdmin.UserHTTP.Lastname,
			Middlename: unitAdmin.UserHTTP.Middlename,
		},
	}

	err := gqlClientWithRequestModifier.Mutate(context.Background(), &m, variables)
	if err != nil {
		log.Fatal("Failed to load data unit admin")
	}
}

func LoadDataForTestRobboUnit() {
	httpClient := &http.Client{}
	gqlClient := graphql.NewClient("http://localhost:8001/query", httpClient)

	robboUnits := []models.NewRobboUnit{
		{
			Name: "name1",
			City: "city1",
		},
		{
			Name: "name2",
			City: "city2",
		},
		{
			Name: "name3",
			City: "city3",
		},
		{
			Name: "name4",
			City: "city4",
		},
		{
			Name: "name5",
			City: "city5",
		},
		{
			Name: "name6",
			City: "city6",
		},
	}

	gqlClientWithRequestModifier := gqlClient.WithRequestModifier(func(request *http.Request) {
		request.Header.Add("Authorization", "Bearer "+viper.GetString("auth.tokens.super_admin"))
	})

	for _, robboUnit := range robboUnits {

		var m struct {
			CreateRobboUnit struct {
				models.RobboUnitHTTP `graphql:"... on RobboUnitHttp"`
				models.Error         `graphql:"... on Error"`
			} `graphql:"CreateRobboUnit(input: $NewRobboUnit)"`
		}

		variables := map[string]interface{}{
			"NewRobboUnit": robboUnit,
		}
		err := gqlClientWithRequestModifier.Mutate(context.Background(), &m, variables)
		if err != nil {
			log.Fatal("Failed to load data robbo units")
		}
	}

	unitAdminRobboUnits := []map[string]interface{}{
		{
			"RobboUnitId": "1",
			"UnitAdminId": "1",
		},
		{
			"RobboUnitId": "2",
			"UnitAdminId": "1",
		},
		{
			"RobboUnitId": "3",
			"UnitAdminId": "1",
		},
		{
			"RobboUnitId": "4",
			"UnitAdminId": "1",
		},
	}

	for _, unitAdminRobboUnit := range unitAdminRobboUnits {

		var m struct {
			SetNewUnitAdminForRobboUnit struct {
				models.Error `graphql:"... on Error"`
			} `graphql:"SetNewUnitAdminForRobboUnit(unitAdminId: $UnitAdminId, robboUnitId: $RobboUnitId)"`
		}

		err := gqlClientWithRequestModifier.Mutate(context.Background(), &m, unitAdminRobboUnit)
		if err != nil {
			log.Fatalf("Failed to load data unitAdmin RobboUnits: %s", err)
		}
	}
}
