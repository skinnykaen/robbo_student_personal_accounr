package factory

import (
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	testmodels "github.com/skinnykaen/robbo_student_personal_account.git/tests/factory/models"
	"github.com/spf13/viper"
)

func TestTableGetAllRobboUnits() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
				},
				"CountRows": 6,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
				},
				"CountRows": 6,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableGetRobboUnitById() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name1",
					City: "city1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name1",
					City: "city1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name1",
					City: "city1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name1",
					City: "city1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Nonexistent robboUnit",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1000",
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableGetRobboUnitsByAccessToken() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
				},
				"CountRows": 4,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
				},
				"CountRows": 4,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"Page":     "1",
				"PageSize": "2",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableGetRobboUnitsByUnitAdminId() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
					{
						ID:   "4",
						Name: "name4",
						City: "city4",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
					{
						ID:   "4",
						Name: "name4",
						City: "city4",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
					{
						ID:   "4",
						Name: "name4",
						City: "city4",
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"UnitAdminId": "1",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableSearchRobboUnitsByName() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
				},
				"CountRows": 6,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{
					{
						ID:   "1",
						Name: "name1",
						City: "city1",
					},
					{
						ID:   "2",
						Name: "name2",
						City: "city2",
					},
					{
						ID:   "3",
						Name: "name3",
						City: "city3",
					},
				},
				"CountRows": 6,
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"Name":     "name",
				"Page":     "1",
				"PageSize": "3",
			},
			Body: map[string]interface{}{
				"RobboUnits": []*models.RobboUnitHTTP{},
				"CountRows":  0,
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableUpdateRobboUnit() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Nonexistent robboUnit",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "100",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableDeleteRobboUnit() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{
					RobboUnitID: "1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "2",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{
					RobboUnitID: "2",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "3",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{
					RobboUnitID: "3",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Nonexistent robboUnit",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "100",
			},
			Body: map[string]interface{}{
				"DeletedRobboUnit": models.DeletedRobboUnit{},
			},
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func TestTableCreateRobboUnit() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "Ok",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "7",
					Name: "name_test",
					City: "city_test",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "8",
					Name: "name_test",
					City: "city_test",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{
					ID:   "9",
					Name: "name_test",
					City: "city_test",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the teacher",
			Token: "Bearer " + viper.GetString("auth.tokens.teacher"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
		{
			Name:  "There is no access to the student",
			Token: "Bearer " + viper.GetString("auth.tokens.student"),
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			Body: map[string]interface{}{
				"RobboUnit": models.RobboUnitHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}
