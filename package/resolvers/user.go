package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input models.NewStudent) (*models.StudentHTTP, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if role < models.UnitAdmin || identityErr != nil {
		return nil, identityErr
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

	studentId, err := r.usersDelegate.CreateStudent(&studentInput, input.ParentID)
	studentInput.UserHTTP.ID = studentId
	return &studentInput, err
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input models.UpdateStudentInput) (*models.StudentHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin, models.Student}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return nil, err
	}
	updateStudentInput := &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			ID:         input.StudentHTTP.UserHTTP.ID,
			Email:      input.StudentHTTP.UserHTTP.Email,
			Firstname:  input.StudentHTTP.UserHTTP.Firstname,
			Lastname:   input.StudentHTTP.UserHTTP.Lastname,
			Middlename: input.StudentHTTP.UserHTTP.Middlename,
			Nickname:   input.StudentHTTP.UserHTTP.Nickname,
		},
	}
	updateStudentErr := r.usersDelegate.UpdateStudent(updateStudentInput)
	if updateStudentErr != nil {
		err := errors.New("baq request")
		return nil, err
	}
	return updateStudentInput, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, studentID string) (string, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return "", err
	}
	_, role, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return "", err
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		err := errors.New("no access")
		return "", err
	}
	id, _ := strconv.ParseUint(studentID, 10, 64)
	deleteStudentErr := r.usersDelegate.DeleteStudent(uint(id))
	if deleteStudentErr != nil {
		err := errors.New("baq request")
		return "", err
	}
	return studentID, nil
}

// GetStudentByAccessToken is the resolver for the GetStudentByAccessToken field.
func (r *queryResolver) GetStudentByAccessToken(ctx context.Context) (*models.StudentHTTP, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return nil, err
	}
	userId, _, identityErr := r.authDelegate.UserIdentity(ginContext)
	if identityErr != nil {
		err := errors.New("status unauthorized")
		return nil, err
	}
	student, err := r.usersDelegate.GetStudentById(userId)
	return student, err
}

// GetStudentByID is the resolver for the GetStudentById field.
func (r *queryResolver) GetStudentByID(ctx context.Context, studentID string) (*models.StudentHTTP, error) {
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
