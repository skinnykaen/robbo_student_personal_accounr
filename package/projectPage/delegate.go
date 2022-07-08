package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage()
	DeleteProjectPage(projectID string) (err error)
	GetProjectPageByID(projectID string) (projectPage *models.ProjectPageHTTP, err error)
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error)
	GetAllProjectPage(authorID string) (projectPage []string, err error)
}
