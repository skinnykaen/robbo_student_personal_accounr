package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateRobboUnit is the resolver for the CreateRobboUnit field.
func (r *mutationResolver) CreateRobboUnit(ctx context.Context, input models.NewRobboUnit) (models.RobboUnitResult, error) {
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

	robboUnitHttp := models.RobboUnitHTTP{
		Name: input.Name,
		City: input.City,
	}

	newRobboUnit, createRobboUnitErr := r.robboUnitsDelegate.CreateRobboUnit(&robboUnitHttp)
	if createRobboUnitErr != nil {
		err := createRobboUnitErr
		return &models.Error{Message: err.Error()}, err
	}
	return &newRobboUnit, nil
}

// UpdateRobboUnit is the resolver for the UpdateRobboUnit field.
func (r *mutationResolver) UpdateRobboUnit(ctx context.Context, input models.UpdateRobboUnit) (models.RobboUnitResult, error) {
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

	updateRobboUnitInput := &models.RobboUnitHTTP{
		ID:   input.ID,
		Name: input.Name,
		City: input.City,
	}

	robboUnitUpdated, updateRobboUnitErr := r.robboUnitsDelegate.UpdateRobboUnit(updateRobboUnitInput)
	if updateRobboUnitErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: updateRobboUnitErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &robboUnitUpdated, nil
}

// DeleteRobboUnit is the resolver for the DeleteRobboUnit field.
func (r *mutationResolver) DeleteRobboUnit(ctx context.Context, robboUnitID string) (*models.DeletedRobboUnit, error) {
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

	deleteRobboUnitErr := r.robboUnitsDelegate.DeleteRobboUnit(robboUnitID)
	if deleteRobboUnitErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: deleteRobboUnitErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.DeletedRobboUnit{RobboUnitID: robboUnitID}, nil
}

// GetRobboUnitByID is the resolver for the GetRobboUnitById field.
func (r *queryResolver) GetRobboUnitByID(ctx context.Context, id string) (models.RobboUnitResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Teacher, models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	robboUnit, getRobboUnitIdErr := r.robboUnitsDelegate.GetRobboUnitById(id)
	if getRobboUnitIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboUnitIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &robboUnit, nil
}

// GetAllRobboUnits is the resolver for the GetAllRobboUnits field.
func (r *queryResolver) GetAllRobboUnits(ctx context.Context, page string, pageSize string) (models.RobboUnitsResult, error) {
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

	robboUnits, countRows, getAllRobboUnitErr := r.robboUnitsDelegate.GetAllRobboUnits(page, pageSize)
	if getAllRobboUnitErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getAllRobboUnitErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// GetRobboUnitsByUnitAdminID is the resolver for the GetRobboUnitsByUnitAdminId field.
func (r *queryResolver) GetRobboUnitsByUnitAdminID(ctx context.Context, unitAdminID string) (models.RobboUnitsResult, error) {
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

	robboUnits, countRows, getRobboUnitsByUnitAdminIdErr := r.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(unitAdminID, "0", "0")
	if getRobboUnitsByUnitAdminIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboUnitsByUnitAdminIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// GetRobboUnitsByAccessToken is the resolver for the GetRobboUnitsByAccessToken field.
func (r *queryResolver) GetRobboUnitsByAccessToken(ctx context.Context, page string, pageSize string) (models.RobboUnitsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	robboUnits, countRows, getRobboUnitsByUnitAdminIdErr := r.robboUnitsDelegate.
		GetRobboUnitsByUnitAdminId(userId, page, pageSize)
	if getRobboUnitsByUnitAdminIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboUnitsByUnitAdminIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}

// SearchRobboUnitsByName is the resolver for the SearchRobboUnitsByName field.
func (r *queryResolver) SearchRobboUnitsByName(ctx context.Context, name string, page string, pageSize string) (models.RobboUnitsResult, error) {
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
	robboUnits, countRows, searchTeachersByEmailErr := r.robboUnitsDelegate.SearchRobboUnitsByName(name, page, pageSize)

	if searchTeachersByEmailErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: searchTeachersByEmailErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboUnitHTTPList{
		RobboUnits: robboUnits,
		CountRows:  countRows,
	}, nil
}
