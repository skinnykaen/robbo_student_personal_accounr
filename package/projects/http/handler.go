package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	authDelegate     auth.Delegate
	projectsDelegate projects.Delegate
}

func NewProjectsHandler(
	authDelegate auth.Delegate,
	projectsDelegate projects.Delegate,
) Handler {
	return Handler{
		authDelegate:     authDelegate,
		projectsDelegate: projectsDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	project := router.Group("/project")
	{
		project.POST("/", h.CreateProject)
		project.GET("/:projectId", h.GetProject)
		project.POST(":projectId", h.UpdateProject)
		project.DELETE("/", h.DeleteProject)
	}
}

type testResponse struct {
	Id string `json:"id"`
}

// TODO нет защиты, так как запрос идет с robboscratchweb без токена. для dev сервера пока можно убрать эти ручки

func (h *Handler) CreateProject(c *gin.Context) {
	fmt.Println("Create Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.Json = string(jsonDataBytes)
	projectId, err := h.projectsDelegate.CreateProject(&projectHTTP)
	fmt.Println(projectId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: projectId,
	})
}

func (h *Handler) GetProject(c *gin.Context) {
	userId, userRole, userIdentityErr := h.authDelegate.UserIdentity(c)
	fmt.Println(userId)
	fmt.Println(userRole)
	if userIdentityErr != nil {
		c.AbortWithError(http.StatusUnauthorized, userIdentityErr)
	}
	allowedRoles := []models.Role{models.Student}
	accessErr := h.authDelegate.UserAccess(userRole, allowedRoles, c)
	if accessErr != nil {
		c.AbortWithError(http.StatusUnauthorized, accessErr)
	}
	projectId := c.Param("projectId")
	if projectId == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	project, err := h.projectsDelegate.GetProjectById(projectId, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(project.Json), &jsonMap)

	c.JSON(http.StatusOK, jsonMap)
}

func (h *Handler) UpdateProject(c *gin.Context) {
	fmt.Println("Update Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	projectId := c.Param("projectId")

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.ID = projectId
	projectHTTP.Json = string(jsonDataBytes)

	err = h.projectsDelegate.UpdateProject(&projectHTTP)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: projectId,
	})
}

func (h *Handler) DeleteProject(c *gin.Context) {

}
