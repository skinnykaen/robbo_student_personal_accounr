package gateway

import (
	"errors"
	"github.com/jackc/pgconn"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UsersGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

func (r *UsersGatewayImpl) GetAllStudents(page, pageSize int, active bool) (
	students []*models.StudentCore,
	countRows int64,
	err error,
) {
	var studentsDB []*models.StudentDB
	offset := (page - 1) * pageSize
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(pageSize).Offset(offset).Where("role = ? AND active = ?", 0, active).
			Find(&studentsDB).Error; err != nil {
			return
		}
		tx.Model(&models.StudentDB{}).Where("active = ?", active).Count(&countRows)
		return
	})

	for _, studentDB := range studentsDB {
		students = append(students, studentDB.ToCore())
	}
	return
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
		if err = tx.Where("email = ? AND  password = ? AND active = true", email, password).
			First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	student = studentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateStudent(student *models.StudentCore) (newStudent *models.StudentCore, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&studentDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return users.ErrAlreadyUsedEmail
		}
		return
	})

	newStudent = studentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteStudent(studentId string) (err error) {
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

func (r *UsersGatewayImpl) UpdateStudent(student *models.StudentCore) (studentUpdated *models.StudentCore, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&studentDb).Clauses(clause.Returning{}).
			Where("id = ?", studentDb.ID).First(&models.StudentDB{}).Updates(
			map[string]interface{}{
				"email":      studentDb.Email,
				"nickname":   studentDb.Nickname,
				"lastname":   studentDb.Lastname,
				"firstname":  studentDb.Firstname,
				"active":     studentDb.Active,
				"middlename": studentDb.Middlename,
			}).Error; err != nil {
			var duplicateEntryError = &pgconn.PgError{Code: "23505"}
			if errors.As(err, &duplicateEntryError) {
				return users.ErrAlreadyUsedEmail
			}
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	studentUpdated = studentDb.ToCore()
	return
}
