package usecase

import (
	"errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
)

type ProjectUseCaseImpl struct {
	projectsGateway    projects.Gateway
	projectPageGateway projectPage.Gateway
}

type ProjectUseCaseModule struct {
	fx.Out
	projects.UseCase
}

func SetupProjectUseCase(projectsGateway projects.Gateway, projectPageGateway projectPage.Gateway) ProjectUseCaseModule {
	return ProjectUseCaseModule{
		UseCase: &ProjectUseCaseImpl{
			projectsGateway:    projectsGateway,
			projectPageGateway: projectPageGateway,
		},
	}
}

func (p *ProjectUseCaseImpl) CreateProject(project *models.ProjectCore) (newProject *models.ProjectCore, err error) {
	return p.projectsGateway.CreateProject(project)
}

func (p *ProjectUseCaseImpl) UpdateProject(project *models.ProjectCore) (err error) {
	return p.projectsGateway.UpdateProject(project)
}

func (p *ProjectUseCaseImpl) DeleteProject() {

}

func (p *ProjectUseCaseImpl) GetProjectById(
	projectId,
	userId string,
	role models.Role,
) (project *models.ProjectCore, err error) {

	project, err = p.projectsGateway.GetProjectById(projectId)
	if err != nil {
		return nil, err
	}
	projectPage, getProjectPageErr := p.projectPageGateway.GetProjectPageByProjectId(projectId)
	if getProjectPageErr != nil {
		return nil, getProjectPageErr
	}
	if projectPage.IsShared == true || role == models.SuperAdmin {
		return project, nil
	}
	if project.AuthorId == userId {
		return
	} else {
		return nil, errors.New("no access")
	}
}
