package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	authDelegate  auth.Delegate
	usersDelegate users.Delegate
}

func NewUsersHandler(
	authDelegate auth.Delegate,
	usersDelegate users.Delegate,
) Handler {
	return Handler{
		authDelegate:  authDelegate,
		usersDelegate: usersDelegate,
	}
}

type getStudentResponse struct {
	Student models.StudentHTTP `json:"student"`
}

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", h.GetUser)
		users.POST("/student", h.CreateStudent)
		users.DELETE("/student/:studentId", h.DeleteStudent)
		users.GET("/student/:studentId", h.GetStudentById)
		users.PUT("/student", h.UpdateStudent)
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	log.Println("Get User")
	userId, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	switch role {
	case models.Student:
		student, getStudentErr := h.usersDelegate.GetStudentById(userId)
		if getStudentErr != nil {
			ErrorHandling(getStudentErr, c)
			return
		}
		c.JSON(http.StatusOK, student)
	case models.Teacher:

		return
	case models.Parent:

		return
	case models.FreeListener:
		return
	case models.UnitAdmin:

		return
	case models.SuperAdmin:

		return
	}
}

func (h *Handler) GetStudentById(c *gin.Context) {
	log.Println("Get Student By Id")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	studentId := c.Param("studentId")

	student, err := h.usersDelegate.GetStudentById(studentId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, getStudentResponse{
		*student,
	})
}

type createStudentInput struct {
	Student  *models.StudentHTTP `json:"student"`
	ParentId string              `json:"parentId"`
}

func (h *Handler) CreateStudent(c *gin.Context) {
	log.Println("Create Student")

	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	input := new(createStudentInput)

	if err := c.BindJSON(input); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	fmt.Println(input.ParentId)

	studentId, err := h.usersDelegate.CreateStudent(input.Student, input.ParentId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"studentId": studentId,
	})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	log.Println("Update Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	userHttp := models.UserHTTP{}

	if err := c.BindJSON(&userHttp); err != nil {
		err = users.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	studentHttp := &models.StudentHTTP{
		UserHTTP: &userHttp,
	}

	err := h.usersDelegate.UpdateStudent(studentHttp)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	log.Println("Delete Student")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	allowedRoles := []models.Role{models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	studentId := c.Param("studentId")
	id, atoiErr := strconv.Atoi(studentId)
	if atoiErr != nil {
		atoiErr = users.ErrBadRequest
		log.Println(atoiErr)
		ErrorHandling(atoiErr, c)
		return
	}
	err := h.usersDelegate.DeleteStudent(uint(id))

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case users.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case users.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case users.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrUserNotFound:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case auth.ErrNotAccess:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
