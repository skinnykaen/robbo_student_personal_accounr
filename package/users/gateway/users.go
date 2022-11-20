package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type UsersGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type UsersGatewayModule struct {
	fx.Out
	users.Gateway
}

func SetupUsersGateway(postgresClient db_client.PostgresClient) UsersGatewayModule {
	return UsersGatewayModule{
		Gateway: &UsersGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *UsersGatewayImpl) AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.StudentDB{}).Where("id = ?", studentId).
			Update("robbo_group_id", gorm.Expr(robboGroupId)).
			Update("robbo_unit_id", gorm.Expr(robboUnitId)).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) GetStudent(email, password string) (student *models.StudentCore, err error) {
	var studentDb models.StudentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	student = studentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateStudent(student *models.StudentCore) (id string, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&studentDb).Error
		return
	})

	id = strconv.FormatUint(uint64(studentDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteStudent(studentId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.StudentDB{}, studentId).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) GetStudentById(studentId string) (student *models.StudentCore, err error) {
	var studentDb models.StudentDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", studentId).First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	student = studentDb.ToCore()

	return
}

func (r *UsersGatewayImpl) UpdateStudent(student *models.StudentCore) (err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&studentDb).Where("id = ?", studentDb.ID).Updates(studentDb).Error
		return
	})
	return
}
