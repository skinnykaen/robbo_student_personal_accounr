package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/skinnykaen/robbo_student_personal_account.git/app/modules"
	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

func NewServer(lifecycle fx.Lifecycle, graphQLModule modules.GraphQLModule, handlers modules.HandlerModule) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := SetupGinRouter(handlers)
				router.Use(TokenAuthMiddleware())
				router.GET("/", playgroundHandler())
				router.POST("/query", graphqlHandler(graphQLModule))

				server := &http.Server{
					Addr: viper.GetString("server.address"),
					Handler: cors.New(
						// TODO make config
						cors.Options{
							AllowedOrigins: []string{
								"http://0.0.0.0:3030",
								"http://0.0.0.0:3000",
								"http://0.0.0.0:8601",
								"http://localhost:3030",
								"http://localhost:3000",
							},
							AllowCredentials: true,
							AllowedMethods: []string{
								http.MethodGet,
								http.MethodPost,
								http.MethodPut,
								http.MethodDelete,
								http.MethodOptions,
								http.MethodOptions,
							},
							AllowedHeaders: []string{"*"},
						},
					).Handler(router),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}

				log.Printf("connect to http://localhost:%s/ for GraphQL playground", viper.GetString("graphqlServer.port"))
				go func() {
					if err = server.ListenAndServe(); err != nil {
						log.Fatalf("Failed to listen and serve: %s", err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}

func SetupGinRouter(handlers modules.HandlerModule) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		GinContextToContextMiddleware(),
	)
	handlers.AuthHandler.InitAuthRoutes(router)
	handlers.ProjectsHandler.InitProjectRoutes(router)
	handlers.ProjectPageHandler.InitProjectRoutes(router)
	handlers.CoursesHandler.InitCourseRoutes(router)
	//handlers.CohortsHandler.InitCohortRoutes(router)
	//handlers.UsersHandler.InitUsersRoutes(router)
	//handlers.RobboUnitsHandler.InitRobboUnitsRoutes(router)
	//handlers.RobboGroupHandler.InitRobboGroupRoutes(router)
	//handlers.CoursePacketHandler.InitCoursePacketRoutes(router)
	return router
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(graphQLModule modules.GraphQLModule) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graphQLModule.UsersResolver,
			},
		))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
