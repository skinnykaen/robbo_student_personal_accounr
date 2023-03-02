package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, peekUserID *string, peekUserRole *int) (models.GetUserResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	var userId string
	var userRole models.Role
	if utils.UseString(peekUserID) == "" || peekUserID == nil {
		userId = ginContext.Value("user_id").(string)
		userRole = ginContext.Value("user_role").(models.Role)
	} else {
		userId = *peekUserID
		userRole = models.Role(*peekUserRole)
	}
	switch userRole {
	case models.Student:
		student, getStudentErr := r.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			return nil, getStudentErr
		}
		return student, nil
	default:
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "internal server error",
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
}
