package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore, parentId string) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	UpdateStudent(student *models.StudentCore) (err error)
}
