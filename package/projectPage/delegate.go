package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage(authorId string) (projectId string, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageById()
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error)
}