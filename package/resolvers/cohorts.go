package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateCohort is the resolver for the CreateCohort field.
func (r *mutationResolver) CreateCohort(ctx context.Context, input models.NewCohort) (models.CohortResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	cohortInput := &models.CohortHTTP{
		Name:           input.Name,
		AssignmentType: input.AssignmentType,
	}

	newCohort, createCohortErr := r.cohortsDelegate.CreateCohort(cohortInput, input.CourseID)
	if createCohortErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: createCohortErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &newCohort, nil
}
