package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateProjectPage(projectPage *models.ProjectPageCore) (id string, err error)
	DeleteProjectPage(projectID string) (err error)
	GetProjectPageByID(projectID string) (projectPage *models.ProjectPageCore, err error)
	UpdateProjectPage(projectPage *models.ProjectPageCore) (err error)
}
