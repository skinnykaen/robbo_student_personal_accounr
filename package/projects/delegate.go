package projects

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateProject(project *models.ProjectHTTP) (newProject *models.ProjectCore, err error)
	DeleteProject()
	GetProjectById(projectId, userId string, role models.Role) (project models.ProjectHTTP, err error)
	UpdateProject(project *models.ProjectHTTP) (err error)
}
