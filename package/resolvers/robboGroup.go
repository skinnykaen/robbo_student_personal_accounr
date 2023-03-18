package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateRobboGroup is the resolver for the CreateRobboGroup field.
func (r *mutationResolver) CreateRobboGroup(ctx context.Context, input models.NewRobboGroup) (models.RobboGroupResult, error) {
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

	robboGroupHttp := models.RobboGroupHTTP{
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	newRobboGroup, createRobboGroupErr := r.robboGroupDelegate.CreateRobboGroup(&robboGroupHttp)
	if createRobboGroupErr != nil {
		err := createRobboGroupErr
		return &models.Error{Message: err.Error()}, err
	}
	return &newRobboGroup, nil
}

// UpdateRobboGroup is the resolver for the UpdateRobboGroup field.
func (r *mutationResolver) UpdateRobboGroup(ctx context.Context, input models.UpdateRobboGroup) (models.RobboGroupResult, error) {
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

	updateRobboGroupInput := &models.RobboGroupHTTP{
		ID:          input.ID,
		Name:        input.Name,
		RobboUnitID: input.RobboUnitID,
	}

	robboGroupUpdated, updateRobboGroupErr := r.robboGroupDelegate.UpdateRobboGroup(updateRobboGroupInput)
	if updateRobboGroupErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: updateRobboGroupErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &robboGroupUpdated, nil
}

// DeleteRobboGroup is the resolver for the DeleteRobboGroup field.
func (r *mutationResolver) DeleteRobboGroup(ctx context.Context, robboGroupID string) (*models.DeletedRobboGroup, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		err := errors.New("internal server error")
		return &models.DeletedRobboGroup{RobboGroupID: ""}, err
	}
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	deleteRobboGroupErr := r.robboGroupDelegate.DeleteRobboGroup(robboGroupID)
	if deleteRobboGroupErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: deleteRobboGroupErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.DeletedRobboGroup{RobboGroupID: robboGroupID}, nil
}

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (models.RobboGroupResult, error) {
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

	robboGroup, getRobboGroupByIdErr := r.robboGroupDelegate.GetRobboGroupById(id)
	if getRobboGroupByIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboGroupByIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &robboGroup, nil
}

// GetRobboGroupsByTeacherID is the resolver for the GetRobboGroupsByTeacherId field.
func (r *queryResolver) GetRobboGroupsByTeacherID(ctx context.Context, teacherID string, page string, pageSize string) (models.RobboGroupsResult, error) {
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

	robboGroups, countRows, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.
		GetRobboGroupsByTeacherId(teacherID, page, pageSize)
	if getRobboGroupsByTeacherIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboGroupsByTeacherIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}

// GetRobboGroupsByRobboUnitID is the resolver for the GetRobboGroupsByRobboUnitId field.
func (r *queryResolver) GetRobboGroupsByRobboUnitID(ctx context.Context, robboUnitID string) (models.RobboGroupsResult, error) {
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

	robboGroups, getRobboGroupsByRobboUnitIdErr := r.robboGroupDelegate.GetRobboGroupsByRobboUnitId(robboUnitID)
	if getRobboGroupsByRobboUnitIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboGroupsByRobboUnitIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
	}, nil
}

// GetRobboGroupsByUnitAdminID is the resolver for the GetRobboGroupsByUnitAdminID field.
func (r *queryResolver) GetRobboGroupsByUnitAdminID(ctx context.Context, unitAdminID string, page string, pageSize string) (models.RobboGroupsResult, error) {
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

	robboGroups, countRows, getRobboGroupsByUnitAdminIdErr := r.robboGroupDelegate.
		GetRobboGroupsByUnitAdminId(unitAdminID, page, pageSize)
	if getRobboGroupsByUnitAdminIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboGroupsByUnitAdminIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}

// GetAllRobboGroupsForUnitAdmin is the resolver for the GetAllRobboGroupsForUnitAdmin field.
func (r *queryResolver) GetAllRobboGroupsForUnitAdmin(ctx context.Context, page string, pageSize string) (models.RobboGroupsResult, error) {
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

	robboGroups, countRows, getAllRobboGroups := r.robboGroupDelegate.
		GetRobboGroupsByUnitAdminId(userId, page, pageSize)
	if getAllRobboGroups != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getAllRobboGroups.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}

// GetAllRobboGroups is the resolver for the GetAllRobboGroups field.
func (r *queryResolver) GetAllRobboGroups(ctx context.Context, page string, pageSize string) (models.RobboGroupsResult, error) {
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

	robboGroups, countRows, getAllRobboGroups := r.robboGroupDelegate.GetAllRobboGroups(page, pageSize)
	if getAllRobboGroups != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getAllRobboGroups.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}

// GetRobboGroupsByAccessToken is the resolver for the GetRobboGroupsByAccessToken field.
func (r *queryResolver) GetRobboGroupsByAccessToken(ctx context.Context, page string, pageSize string) (models.RobboGroupsResult, error) {
	ginContext, getGinContextErr := GinContextFromContext(ctx)
	if getGinContextErr != nil {
		return nil, getGinContextErr
	}
	userId := ginContext.Value("user_id").(string)
	userRole := ginContext.Value("user_role").(models.Role)
	allowedRoles := []models.Role{models.Teacher, models.SuperAdmin, models.UnitAdmin}
	accessErr := r.authDelegate.UserAccess(userRole, allowedRoles, ctx)
	if accessErr != nil {
		return nil, accessErr
	}

	robboGroups, countRows, getRobboGroupsByTeacherIdErr := r.robboGroupDelegate.
		GetRobboGroupsByTeacherId(userId, page, pageSize)
	if getRobboGroupsByTeacherIdErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: getRobboGroupsByTeacherIdErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}

// SearchGroupsByName is the resolver for the SearchGroupsByName field.
func (r *queryResolver) SearchGroupsByName(ctx context.Context, name string, page string, pageSize string) (models.RobboGroupsResult, error) {
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
	robboGroups, countRows, searchRobboGroupByNameErr := r.robboGroupDelegate.SearchRobboGroupByName(name, page, pageSize)

	if searchRobboGroupByNameErr != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: searchRobboGroupByNameErr.Error(),
			Extensions: map[string]interface{}{
				"code": "500",
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: robboGroups,
		CountRows:   countRows,
	}, nil
}
