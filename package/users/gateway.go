package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetStudent(email, password string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	UpdateStudent(student *models.StudentCore) (err error)
}
