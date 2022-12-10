package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
)

type UsersUseCaseImpl struct {
	Gateway users.Gateway
}

type UsersUseCaseModule struct {
	fx.Out
	users.UseCase
}

func SetupUsersUseCase(gateway users.Gateway) UsersUseCaseModule {
	return UsersUseCaseModule{
		UseCase: &UsersUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *UsersUseCaseImpl) GetStudentById(studentId string) (student *models.StudentCore, err error) {
	return p.Gateway.GetStudentById(studentId)
}

func (p *UsersUseCaseImpl) CreateStudent(student *models.StudentCore, parentId string) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(student.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	student.Password = passwordHash
	id, err = p.Gateway.CreateStudent(student)
	if err != nil {
		return
	}
	return
}

func (p *UsersUseCaseImpl) DeleteStudent(studentId uint) (err error) {
	return p.Gateway.DeleteStudent(studentId)
}

func (p *UsersUseCaseImpl) UpdateStudent(student *models.StudentCore) (updatedStudent *models.StudentCore, err error) {
	updatedStudent, err = p.Gateway.UpdateStudent(student)
	if err != nil {
		log.Println("Error update student")
		return
	}
	return
}
