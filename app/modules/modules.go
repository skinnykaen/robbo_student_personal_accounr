package modules

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	authdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate"
	authgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	authusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	crsdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/delegate"
	crsgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/gateway"
	crshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http"
	crsusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	edxusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/edx/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	ppagedelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/delegate"
	ppagegateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/gateway"
	ppagehttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/http"
	ppageusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	prjdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/delegate"
	prjgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/gateway"
	prjhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	prjusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/resolvers"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	usersdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/users/delegate"
	usersgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/users/gateway"
	usershtpp "github.com/skinnykaen/robbo_student_personal_account.git/package/users/http"
	usersusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/users/usecase"
)

type GatewayModule struct {
	AuthGateway        auth.Gateway
	CoursesGateway     courses.Gateway
	ProjectPageGateway projectPage.Gateway
	ProjectsGateway    projects.Gateway
	UsersGateway       users.Gateway
}

func SetupGateway(postgresClient db_client.PostgresClient) GatewayModule {
	return GatewayModule{
		AuthGateway:        authgateway.SetupAuthGateway(postgresClient),
		CoursesGateway:     crsgateway.SetupCoursesGateway(postgresClient),
		ProjectPageGateway: ppagegateway.SetupProjectPageGateway(postgresClient),
		ProjectsGateway:    prjgateway.SetupProjectsGateway(postgresClient),
		UsersGateway:       usersgateway.SetupUsersGateway(postgresClient),
	}
}

type UseCaseModule struct {
	AuthUseCase        auth.UseCase
	CoursesUseCase     courses.UseCase
	EdxUseCase         edx.UseCase
	ProjectPageUseCase projectPage.UseCase
	ProjectsUseCase    projects.UseCase
	UsersUseCase       users.UseCase
}

func SetupUseCase(gateway GatewayModule) UseCaseModule {
	return UseCaseModule{
		AuthUseCase:        authusecase.SetupAuthUseCase(gateway.UsersGateway),
		CoursesUseCase:     crsusecase.SetupCourseUseCase(gateway.CoursesGateway),
		EdxUseCase:         edxusecase.SetupEdxApiUseCase(),
		ProjectPageUseCase: ppageusecase.SetupProjectPageUseCase(gateway.ProjectPageGateway, gateway.ProjectsGateway),
		ProjectsUseCase:    prjusecase.SetupProjectUseCase(gateway.ProjectsGateway, gateway.ProjectPageGateway),
		UsersUseCase:       usersusecase.SetupUsersUseCase(gateway.UsersGateway),
	}
}

type DelegateModule struct {
	AuthDelegate        auth.Delegate
	CoursesDelegate     courses.Delegate
	ProjectPageDelegate projectPage.Delegate
	ProjectsDelegate    projects.Delegate
	UsersDelegate       users.Delegate
}

func SetupDelegate(usecase UseCaseModule) DelegateModule {
	return DelegateModule{
		AuthDelegate:        authdelegate.SetupAuthDelegate(usecase.AuthUseCase),
		CoursesDelegate:     crsdelegate.SetupCourseDelegate(usecase.CoursesUseCase, usecase.EdxUseCase),
		ProjectPageDelegate: ppagedelegate.SetupProjectPageDelegate(usecase.ProjectPageUseCase),
		ProjectsDelegate:    prjdelegate.SetupProjectDelegate(usecase.ProjectsUseCase),
		UsersDelegate:       usersdelegate.SetupUsersDelegate(usecase.UsersUseCase),
	}
}

type HandlerModule struct {
	ProjectsHandler    prjhttp.Handler
	ProjectPageHandler ppagehttp.Handler
	AuthHandler        authhttp.Handler
	CoursesHandler     crshttp.Handler
	UsersHandler       usershtpp.Handler
}

func SetupHandler(delegate DelegateModule) HandlerModule {
	return HandlerModule{
		ProjectsHandler:    prjhttp.NewProjectsHandler(delegate.AuthDelegate, delegate.ProjectsDelegate),
		ProjectPageHandler: ppagehttp.NewProjectPageHandler(delegate.AuthDelegate, delegate.ProjectsDelegate, delegate.ProjectPageDelegate),
		AuthHandler:        authhttp.NewAuthHandler(delegate.AuthDelegate),
		CoursesHandler:     crshttp.NewCoursesHandler(delegate.AuthDelegate, delegate.CoursesDelegate),
		UsersHandler:       usershtpp.NewUsersHandler(delegate.AuthDelegate, delegate.UsersDelegate),
	}
}

type GraphQLModule struct {
	UsersResolver resolvers.Resolver
}

func SetupGraphQLModule(delegate DelegateModule) GraphQLModule {
	return GraphQLModule{
		UsersResolver: resolvers.NewResolver(
			delegate.AuthDelegate,
			delegate.UsersDelegate,
			delegate.CoursesDelegate,
			delegate.ProjectPageDelegate,
		),
	}
}
