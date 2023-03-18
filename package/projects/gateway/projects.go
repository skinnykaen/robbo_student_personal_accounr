package gateway

import (
	"errors"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"strconv"
)

type ProjectsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type ProjectsGatewayModule struct {
	fx.Out
	projects.Gateway
}

func SetupProjectsGateway(postgresClient db_client.PostgresClient) ProjectsGatewayModule {
	return ProjectsGatewayModule{
		Gateway: &ProjectsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *ProjectsGatewayImpl) CreateProject(project *models.ProjectCore) (id string, err error) {
	projectDb := models.ProjectDB{}
	projectDb.FromCore(project)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&projectDb).Error
		return
	})

	id = strconv.FormatUint(uint64(projectDb.ID), 10)
	return
}

func (r *ProjectsGatewayImpl) GetProjectById(projectId, userId string) (project *models.ProjectCore, err error) {
	var projectDb models.ProjectDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", projectId).First(&projectDb).Error; err != nil {
			return
		}
		return
	})

	project = projectDb.ToCore()
	if project.AuthorId == userId {
		return
	} else {
		return nil, errors.New("NO access")
	}
}

func (r *ProjectsGatewayImpl) GetProjectsByAuthorId(authorId string, page, pageSize int) (
	projects []*models.ProjectCore,
	countRows int64,
	err error,
) {
	var projectsDb []*models.ProjectDB
	offset := (page - 1) * pageSize
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		fmt.Println(authorId)
		if err = tx.Limit(pageSize).Offset(offset).Where("author_id = ?", authorId).
			Find(&projectsDb).Error; err != nil {
			return
		}
		tx.Model(&models.ProjectDB{}).Where("author_id = ?", authorId).Count(&countRows)
		return
	})

	for _, projectDb := range projectsDb {
		projects = append(projects, projectDb.ToCore())
	}
	return
}

func (r *ProjectsGatewayImpl) DeleteProject(projectId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.ProjectDB{}, projectId).Error
		return
	})
	return
}

func (r *ProjectsGatewayImpl) UpdateProject(project *models.ProjectCore) (err error) {
	projectDb := models.ProjectDB{}
	projectDb.FromCore(project)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&projectDb).Where("ID = ?", projectDb.ID).Updates(projectDb).Error
		return
	})
	return
}
