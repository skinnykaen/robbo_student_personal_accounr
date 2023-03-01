package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetStudent(email, password string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore) (newStudent *models.StudentCore, err error)
	DeleteStudent(studentId string) (err error)
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	UpdateStudent(student *models.StudentCore) (studentUpdated *models.StudentCore, err error)
}
