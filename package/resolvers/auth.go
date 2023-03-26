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

// SignIn is the resolver for the SignIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input models.SignInInput) (models.SignInResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	accessToken, refreshToken, err := r.authDelegate.SignIn(input.Email, input.Password, uint(input.UserRole))
	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	utils.SetRefreshToken(refreshToken, ginContext)
	return &models.SignInResponse{
		AccessToken: accessToken,
	}, nil
}

// SignUpRequest is the resolver for the SignUpRequest field.
func (r *mutationResolver) SignUpRequest(ctx context.Context, input models.NewStudent) (*models.Error, error) {
	studentInput := models.StudentHTTP{
		UserHTTP: &models.UserHTTP{
			Email:      input.Email,
			Password:   input.Password,
			Firstname:  input.Firstname,
			Lastname:   input.Lastname,
			Middlename: input.Middlename,
			Nickname:   input.Nickname,
			Active:     false,
			Role:       0,
		},
	}
	_, createStudentErr := r.usersDelegate.CreateStudent(&studentInput, "")
	if createStudentErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: createStudentErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.Error{}, nil
}

// SignOut is the resolver for the SignOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (*models.Error, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	utils.SetRefreshToken("", ginContext)
	return &models.Error{}, nil
}

// Refresh is the resolver for the Refresh field.
func (r *mutationResolver) Refresh(ctx context.Context) (models.SignInResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	refreshToken, err := utils.GetRefreshToken(ginContext)
	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": "401",
			},
		}
	}

	newAccessToken, err := r.authDelegate.RefreshToken(refreshToken)
	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.SignInResponse{
		AccessToken: newAccessToken,
	}, nil
}
