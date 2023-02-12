package factory

import (
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	testmodels "github.com/skinnykaen/robbo_student_personal_account.git/tests/factory/models"
	"github.com/spf13/viper"
)

func TestTableCreateCourseAccessRelationRobboGroup() (data []testmodels.TestTable) {
	data = []testmodels.TestTable{
		{
			Name:  "There is no access without a token",
			Token: "",
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{},
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
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{
					ID:        "1",
					Parameter: "robbo_group",
					CourseID:  "1",
					ObjectID:  "1",
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
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{
					ID:        "2",
					Parameter: "robbo_group",
					CourseID:  "1",
					ObjectID:  "1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is access to the unit admin",
			Token: "Bearer " + viper.GetString("auth.tokens.unit_admin"),
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{
					ID:        "3",
					Parameter: "robbo_group",
					CourseID:  "1",
					ObjectID:  "1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name:  "There is no access to the free listener",
			Token: "Bearer " + viper.GetString("auth.tokens.free_listener"),
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{},
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
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{},
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
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{},
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
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			Body: map[string]interface{}{
				"CourseRelation": models.CourseRelationHTTP{},
			},
			ExpectedError: graphql.Error{
				Message:   auth.ErrNotAccess.Error(),
				Locations: nil,
			},
		},
	}
	return
}
