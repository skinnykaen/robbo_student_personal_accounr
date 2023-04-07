package projects

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateProject(project *models.ProjectCore) (newProject *models.ProjectCore, err error)
	DeleteProject()
	GetProjectById(projectId, userId string, role models.Role) (project *models.ProjectCore, err error)
	UpdateProject(project *models.ProjectCore) (err error)
}
