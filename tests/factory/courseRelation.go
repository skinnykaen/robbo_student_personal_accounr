package factory

import (
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
)

func DataCreateCourseAccessRelationRobboGroup() (data []models.DataTest) {
	data = []models.DataTest{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
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
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the super admin",
			Token: "Bearer " + viper.GetString("auth.tokens.super_admin"),
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			ExpectedError: nil,
		},
	}
	return
}
