package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input models.NewStudent) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	studentInput := models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       0,
		},
	}

	newStudent, createStudentErr := r.usersDelegate.CreateStudent(&studentInput, input.ParentID)
	if createStudentErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: createStudentErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return newStudent, nil
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input models.UpdateProfileInput) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Student}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	updateStudentInput := &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.ID,
			Email:      input.Email,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Role:       0,
		},
	}

	studentUpdated, updateStudentErr := r.usersDelegate.UpdateStudent(updateStudentInput)
	if updateStudentErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: updateStudentErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return studentUpdated, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (*models.DeletedStudent, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	deleteStudentErr := r.usersDelegate.DeleteStudent(studentID)
	if deleteStudentErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: deleteStudentErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.DeletedStudent{StudentID: studentID}, nil
}

// GetStudentByAccessToken is the resolver for the GetStudentByAccessToken field.
func (r *queryResolver) GetStudentByAccessToken(ctx context.Context) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	student, err := r.usersDelegate.GetStudentById(userId)
	return student, err
}

// GetStudentByID is the resolver for the GetStudentById field.
func (r *queryResolver) GetStudentByID(ctx context.Context, studentID string) (models.StudentResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	student, err := r.usersDelegate.GetStudentById(studentID)
	return student, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
