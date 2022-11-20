package users

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	CreateStudent(student *models.StudentHTTP, parentId string) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	GetStudentById(studentId string) (student *models.StudentHTTP, err error)
	UpdateStudent(student *models.StudentHTTP) (err error)
}
