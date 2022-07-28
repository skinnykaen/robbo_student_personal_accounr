package users

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	//GetStudent(email, password string) (student models.StudentHTTP, err error)
	CreateStudent(student *models.StudentHTTP) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	GetStudentById(studentId uint) (student models.StudentHTTP, err error)
	UpdateStudent(student *models.StudentHTTP) (err error)

	//GetTeacher(email, password string) (teacher models.TeacherHTTP, err error)
	GetTeacherById(teacherId uint) (teacher models.TeacherHTTP, err error)
	CreateTeacher(teacher *models.TeacherHTTP) (id string, err error)
	UpdateTeacher(teacherHTTP *models.TeacherHTTP) (err error)
	DeleteTeacher(teacherId uint) (err error)

	CreateParent(parentHTTP *models.ParentHTTP) (id string, err error)
	UpdateParent(parentHTTP *models.ParentHTTP) (err error)
	DeleteParent(parentId uint) (err error)
	GetParentById(parentId uint) (parent models.ParentHTTP, err error)
	GetAllParent() (parents []*models.ParentHTTP, err error)
	//GetParent(email, password string) (parent models.ParentHTTP, err error)

	CreateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (id string, err error)
	UpdateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (err error)
	DeleteFreeListener(freeListenerId uint) (err error)
	GetFreeListenerById(freeListenerId uint) (freeListener models.FreeListenerHttp, err error)
	//GetFreeListener(email, password string) (freeListener models.FreeListenerHttp, err error)

	CreateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (id string, err error)
	UpdateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (err error)
	DeleteUnitAdmin(unitAdminId uint) (err error)
	GetUnitAdminById(unitAdminId uint) (unitAdmin models.UnitAdminHTTP, err error)
	//GetUnitAdmin(email, password string) (unitAdmin models.UnitAdminHTTP, err error)

	//GetSuperAdmin(email, password string) (superAdmin models.SuperAdminHTTP, err error)
	GetSuperAdminById(superAdminId uint) (superAdmin models.SuperAdminHTTP, err error)
	UpdateSuperAdmin(superAdminHTTP *models.SuperAdminHTTP) (err error)
	DeleteSuperAdmin(superAdminId uint) (err error)
}
