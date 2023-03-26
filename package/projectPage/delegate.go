package projectPage

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProjectPage(authorId string) (newProjectPage models.ProjectPageHTTP, err error)
	DeleteProjectPage(projectId string) (err error)
	GetProjectPageById(projectPageId string) (projectPage models.ProjectPageHTTP, err error)
	GetAllProjectPagesByUserId(authorId, page, pageSize string) (
		projectPages []*models.ProjectPageHTTP,
		countRows int,
		err error,
	)
	GetAllProjectPages(page, pageSize string) (projectPages []*models.ProjectPageHTTP, countRows int, err error)
	UpdateProjectPage(projectPage *models.ProjectPageHTTP) (projectPageUpdated models.ProjectPageHTTP, err error)
}
