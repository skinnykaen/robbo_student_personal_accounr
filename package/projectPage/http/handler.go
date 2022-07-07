package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"net/http"
)

type Handler struct {
	authDelegate        auth.Delegate
	projectsDelegate    projects.Delegate
	projectPageDelegate projectPage.Delegate
}

type updateInput struct {
	ProjectPage *models.ProjectPageHTTP `json:"project_page"`
}

func NewProjectPageHandler(authDelegate auth.Delegate, projectsDelegate projects.Delegate, projectPageDelegate projectPage.Delegate) Handler {
	return Handler{
		authDelegate:        authDelegate,
		projectsDelegate:    projectsDelegate,
		projectPageDelegate: projectPageDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	projectPage := router.Group("/projectPage")
	{
		projectPage.POST("/", h.CreateProjectPage)
		projectPage.GET("/:projectPageId", h.GetProjectPageByID)
		projectPage.PUT("/", h.UpdateProjectPage)
		projectPage.DELETE("/:projectPageID", h.DeleteProjectPage)
	}
}

type testResponse struct {
	Id string `json:"id"`
}

type testRequest struct {
	Body string `json:"body"`
}

func (h *Handler) CreateProjectPage(c *gin.Context) {
}

func (h *Handler) GetProjectPageByID(c *gin.Context) {
	fmt.Println("Get Project Page By ID")
	projectID := c.Param("projectId")
	projectPageByID, err := h.projectPageDelegate.GetProjectPageByID(projectID)
	if err == nil {

		return
	} else {
		switch err {
		case projectPage.ErrPageNotFound:
			c.AbortWithStatus(http.StatusNotFound)
			return
		case projectPage.ErrInternalServerLevel:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case projectPage.ErrBadRequest:
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

func (h *Handler) UpdateProjectPage(c *gin.Context) {
	fmt.Println("Update Project Page")
	inp := new(updateInput)
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.projectPageDelegate.UpdateProjectPage(inp.ProjectPage)
	if err == nil {
		c.AbortWithStatus(http.StatusOK)
		return
	} else {
		switch err {
		case projectPage.ErrPageNotFound:
			c.AbortWithStatus(http.StatusNotFound)
			return
		case projectPage.ErrInternalServerLevel:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case projectPage.ErrBadRequest:
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

}
func (h *Handler) DeleteProjectPage(c *gin.Context) {
	fmt.Println("Delete Project Page")

	projectID := c.Param("projectId")
	err := h.projectPageDelegate.DeleteProjectPage(projectID)
	if err == nil {
		c.AbortWithStatus(http.StatusOK)
		return
	} else {
		switch err {
		case projectPage.ErrPageNotFound:
			c.AbortWithStatus(http.StatusNotFound)
			return
		case projectPage.ErrInternalServerLevel:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case projectPage.ErrBadRequest:
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

}
