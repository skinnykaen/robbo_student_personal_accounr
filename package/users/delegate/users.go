package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersDelegateImpl struct {
	UseCase users.UseCase
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

func (p *UsersDelegateImpl) CreateStudent(student *models.StudentHTTP, parentId string) (id string, err error) {
	studentCore := student.ToCore()
	return p.UseCase.CreateStudent(studentCore, parentId)
}

func (p *UsersDelegateImpl) DeleteStudent(studentId uint) (err error) {
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

func (p *UsersDelegateImpl) UpdateStudent(studentHTTP *models.StudentHTTP) (err error) {
	studentCore := studentHTTP.ToCore()
	return p.UseCase.UpdateStudent(studentCore)
}
