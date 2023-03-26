package users

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	CreateStudent(student *models.StudentHTTP, parentId string) (newStudent *models.StudentHTTP, err error)
	DeleteStudent(studentId string) (err error)
	GetStudentById(studentId string) (student *models.StudentHTTP, err error)
	GetAllStudents(page, pageSize string, active bool) (students []*models.StudentHTTP, countRows int, err error)
	SetActiveForStudent(studentId string, active bool) (err error)
	UpdateStudent(student *models.StudentHTTP) (updatedStudent models.StudentHTTP, err error)
}
