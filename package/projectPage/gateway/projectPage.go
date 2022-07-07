package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"strconv"
)

type ProjectPageGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type ProjectPageGatewayModule struct {
	fx.Out
	projectPage.Gateway
}

func SetupProjectPageGateway(postgresClient db_client.PostgresClient) ProjectPageGatewayModule {
	return ProjectPageGatewayModule{
		Gateway: &ProjectPageGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *ProjectPageGatewayImpl) CreateProjectPage(projectPage *models.ProjectPageCore) (id string, err error) {
	projectPageDb := models.ProjectPageDB{}
	projectPageDb.FromCore(projectPage)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&projectPageDb).Error
		return
	})

	id = strconv.FormatUint(uint64(projectPageDb.ID), 10)
	return
}

func (r *ProjectPageGatewayImpl) GetProjectPageByID(projectID string) (projectPage *models.ProjectPageCore, err error) {
	var projectPageDB models.ProjectPageDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", projectID).First(&projectPageDB).Error; err != nil {
			return
		}
		return
	})

	projectPage = projectPageDB.ToCore()

	return
}

func (r *ProjectPageGatewayImpl) DeleteProjectPage(projectID string) (err error) {
	id, _ := strconv.Atoi(projectID)
	projectPageDb := models.ProjectPageDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&projectPageDb).Delete(projectPageDb, id).Error
		return
	})
	return
}

func (r *ProjectPageGatewayImpl) UpdateProjectPage(projectPage *models.ProjectPageCore) (err error) {
	projectPageDb := models.ProjectPageDB{}
	projectPageDb.FromCore(projectPage)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&projectPageDb).Where("ID = ?", projectPageDb.ID).Updates(projectPageDb).Error
		return
	})
	return
}
