package factory

import (
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"github.com/spf13/viper"
)

func DataCreateRobboUnit() (data []models.DataTest) {
	data = []models.DataTest{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
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
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func DataUpdateRobboUnit() (data []models.DataTest) {

	data = []models.DataTest{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"UpdateRobboUnit": models.UpdateRobboUnit{
					ID:   "1",
					Name: "name_test",
					City: "city_test",
				},
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
					ID:   "1000",
					Name: "name_updated",
					City: "city_updated",
				},
			},
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func DataDeleteRobboUnit() (data []models.DataTest) {
	data = []models.DataTest{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
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
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "2",
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "3",
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"RobboUnitId": "3",
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
				"RobboUnitId": "3",
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
				"RobboUnitId": "3",
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
				"RobboUnitId": "3",
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
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}

func DataGetRobboUnitById() (data []models.DataTest) {
	data = []models.DataTest{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
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
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"RobboUnitId": "1",
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"RobboUnitId": "3",
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
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the parent",
			Token: "Bearer " + viper.GetString("auth.tokens.parent"),
			Variables: map[string]interface{}{
				"RobboUnitId": "3",
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
				"RobboUnitId": "3",
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
			ExpectedError: graphql.Error{
				Message:   robboUnits.ErrRobboUnitNotFound.Error(),
				Locations: nil,
			},
		},
	}
	return
}
