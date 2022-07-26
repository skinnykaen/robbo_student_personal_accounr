package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateParent(parent *models.ParentCore) (id string, err error)
	GetParentById(parentId uint) (parent *models.ParentCore, err error)
	GetUnitAdminById(unitAdminId uint) (unitAdmin *models.UnitAdminCore, err error)
	GetTeacherById(teacherId uint) (teacher *models.TeacherCore, err error)
	UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error)
	CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error)
	DeleteUnitAdmin(unitAdminId uint) (err error)
	CreateTeacher(teacher *models.TeacherCore) (id string, err error)
	UpdateTeacher(teacher *models.TeacherCore) (err error)
	DeleteTeacher(teacherId uint) (err error)
	UpdateParent(parent *models.ParentCore) (err error)
	DeleteParent(parentId uint) (err error)
	GetTeacher(email, password string) (teacher *models.TeacherCore, err error)
	GetParent(email, password string) (parent *models.ParentCore, err error)
	GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error)
	GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error)
	GetStudent(email, password string) (student *models.StudentCore, err error)
	CreateStudent(student *models.StudentCore) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	GetStudentById(studentId uint) (student *models.StudentCore, err error)
	UpdateStudent(student *models.StudentCore) (err error)
	GetSuperAdminById(superAdminId uint) (superAdmin *models.SuperAdminCore, err error)
}
