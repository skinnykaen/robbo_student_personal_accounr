package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// SingIn is the resolver for the SingIn field.
func (r *mutationResolver) SingIn(ctx context.Context, input models.SignInInput) (models.SignInResult, error) {
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
	setRefreshToken(refreshToken, ginContext)
	return &models.SingInResponse{
		AccessToken: accessToken,
	}, nil
}

// SingOut is the resolver for the SingOut field.
func (r *mutationResolver) SingOut(ctx context.Context) (*models.Error, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	setRefreshToken("", ginContext)
	return &models.Error{}, nil
}

// RequestResetPassword is the resolver for the RequestResetPassword field.
func (r *mutationResolver) RequestResetPassword(ctx context.Context, email string) (*models.Error, error) {
	err := r.authDelegate.RequestResetPassword(email)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// ConfirmResetPassword is the resolver for the ConfirmResetPassword field.
func (r *mutationResolver) ConfirmResetPassword(ctx context.Context, email string, verifyCode string) (*models.Error, error) {
	err := r.authDelegate.ConfirmResetPassword(email, verifyCode)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// Refresh is the resolver for the Refresh field.
func (r *mutationResolver) Refresh(ctx context.Context) (models.SignInResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	refreshToken, err := getRefreshToken(ginContext)
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
	return &models.SingInResponse{
		AccessToken: newAccessToken,
	}, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func getRefreshToken(c *gin.Context) (refreshToken string, err error) {
	refreshToken = c.Value("refresh_token").(string)
	if refreshToken == "" {
		return "", errors.New("error finding cookie")
	}
	return
}
func setRefreshToken(value string, c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		value,
		60*60*24*7,
		"/",
		"0.0.0.0",
		false,
		false,
	)
}
