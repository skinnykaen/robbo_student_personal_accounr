package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type UsersDelegateImpl struct {
	UseCase users.UseCase
}

func (p *UsersDelegateImpl) SetActiveForStudent(studentId string, active bool) (err error) {
	return p.UseCase.SetActiveForStudent(studentId, active)
}

func (p *UsersDelegateImpl) GetAllStudents(page, pageSize string, active bool) (
	students []*models.StudentHTTP,
	countRows int,
	err error,
) {
	pageInt32, _ := strconv.ParseInt(page, 10, 32)
	pageSizeInt32, _ := strconv.ParseInt(pageSize, 10, 32)
	studentsCore, countRowsInt64, err := p.UseCase.GetAllStudents(int(pageInt32), int(pageSizeInt32), active)
	if err != nil {
		return
	}
	countRows = int(countRowsInt64)
	for _, studentCore := range studentsCore {
		studentTemp := models.StudentHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		studentTemp.FromCore(studentCore)
		students = append(students, &studentTemp)
	}
	return
}

type UsersDelegateModule struct {
	fx.Out
	users.Delegate
}

func SetupUsersDelegate(usecase users.UseCase) UsersDelegateModule {
	return UsersDelegateModule{
		Delegate: &UsersDelegateImpl{
			usecase,
		},
	}
}

func (p *UsersDelegateImpl) CreateStudent(
	student *models.StudentHTTP,
	parentId string,
) (newStudent *models.StudentHTTP, err error) {
	studentCore := student.ToCore()
	newStudentCore, err := p.UseCase.CreateStudent(studentCore, parentId)
	if err != nil {
		log.Println(err)
		return
	}
	newStudent = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	newStudent.FromCore(newStudentCore)
	return
}

func (p *UsersDelegateImpl) DeleteStudent(studentId string) (err error) {
	return p.UseCase.DeleteStudent(studentId)
}

func (p *UsersDelegateImpl) GetStudentById(studentId string) (student *models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.GetStudentById(studentId)
	if err != nil {
		log.Println("User not found")
		return student, auth.ErrUserNotFound
	}
	student = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	student.FromCore(studentCore)
	return
}

func (p *UsersDelegateImpl) UpdateStudent(studentHTTP *models.StudentHTTP) (updatedStudent models.StudentHTTP, err error) {
	studentCore := studentHTTP.ToCore()
	updatedStudentCore, err := p.UseCase.UpdateStudent(studentCore)
	if err != nil {
		return
	}
	updatedStudent = models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	updatedStudent.FromCore(updatedStudentCore)
	return
}
