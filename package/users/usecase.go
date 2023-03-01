package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore, parentId string) (newStudent *models.StudentCore, err error)
	DeleteStudent(studentId string) (err error)
	UpdateStudent(student *models.StudentCore) (updatedStudent *models.StudentCore, err error)
}
