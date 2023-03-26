package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/utils"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type UsersUseCaseImpl struct {
	Gateway users.Gateway
}

func (p *UsersUseCaseImpl) SetActiveForStudent(studentId string, active bool) (err error) {
	student, getStudentErr := p.Gateway.GetStudentById(studentId)
	if getStudentErr != nil {
		return getStudentErr
	}
	if active == true {
		student.Active = true
		if _, updateErr := p.Gateway.UpdateStudent(student); updateErr != nil {
			return updateErr
		}
		subject := "Ваш профиль был активирован"
		body := "<p> Профиль " + student.Email + " был успешно активирован администратором!</p>"
		if err = utils.SendEmail(subject, student.Email, body); err != nil {
			return err
		}
		return nil
	} else {
		student.Active = false
		if _, updateErr := p.Gateway.UpdateStudent(student); updateErr != nil {
			return updateErr
		}
		fmt.Println(student)
		subject := "Ваш профиль был деактивирован"
		body := "<p> Профиль " + student.Email + " был деактивирован по решению администратора.</p>"
		if err = utils.SendEmail(subject, student.Email, body); err != nil {
			return err
		}
		return nil
	}
}

func (p *UsersUseCaseImpl) GetAllStudents(page, pageSize int, active bool) (students []*models.StudentCore, countRows int64, err error) {
	return p.Gateway.GetAllStudents(page, pageSize, active)
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

func (p *UsersUseCaseImpl) CreateStudent(student *models.StudentCore, parentId string) (newStudent *models.StudentCore, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(student.Password))
	pwd.Write([]byte(viper.GetString("auth_hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	student.Password = passwordHash
	newStudent, err = p.Gateway.CreateStudent(student)
	if err != nil {
		return
	}
	return
}

func (p *UsersUseCaseImpl) DeleteStudent(studentId string) (err error) {
	return p.Gateway.DeleteStudent(studentId)
}

func (p *UsersUseCaseImpl) UpdateStudent(student *models.StudentCore) (studentUpdated *models.StudentCore, err error) {
	studentUpdated, err = p.Gateway.UpdateStudent(student)
	if err != nil {
		return
	}
	return
}
