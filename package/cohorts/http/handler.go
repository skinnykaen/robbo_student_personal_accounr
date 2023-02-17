package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate    auth.Delegate
	cohortsDelegate cohorts.Delegate
}

func NewCohortsHandler(
	authDelegate auth.Delegate,
	cohortsDelegate cohorts.Delegate,
) Handler {
	return Handler{
		authDelegate:    authDelegate,
		cohortsDelegate: cohortsDelegate,
	}
}

type testCohortResponse struct {
	CohortID string `json:"cohortId"`
}

func (h *Handler) InitCohortRoutes(router *gin.Engine) {
	cohort := router.Group("/cohort")
	{
		cohort.POST("/createCohort/:courseId", h.CreateCohort)
		cohort.POST("/addStudent/:courseId/:cohortId/:studentId", h.AddStudent)
	}
}

func (h *Handler) CreateCohort(c *gin.Context) {
	fmt.Println("Create Cohort")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles, c)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	createCohortResponse := models.CohortHTTP{}
	courseId := c.Param("courseId")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		err = cohorts.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	err = json.Unmarshal(body, &createCohortResponse)
	fmt.Println(createCohortResponse)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	cohortHTTP := models.CohortHTTP{}

	_, err = h.cohortsDelegate.CreateCohort(&createCohortResponse, courseId)

	fmt.Println(cohortHTTP)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) AddStudent(c *gin.Context) {
	fmt.Println("Add Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles, c)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	cohortId := c.Param("cohortId")
	courseId := c.Param("courseId")
	studentId := c.Param("studentId")
	err := h.cohortsDelegate.AddStudentToCohort(courseId, cohortId, studentId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case cohorts.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case cohorts.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case cohorts.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrNotAccess:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
