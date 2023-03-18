package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"go.uber.org/fx"
)

type ProjectUseCaseImpl struct {
	projects.Gateway
}

type ProjectUseCaseModule struct {
	fx.Out
	projects.UseCase
}

func SetupProjectUseCase(gateway projects.Gateway) ProjectUseCaseModule {
	return ProjectUseCaseModule{
		UseCase: &ProjectUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *ProjectUseCaseImpl) CreateProject(project *models.ProjectCore) (id string, err error) {
	return p.Gateway.CreateProject(project)
}

func (p *ProjectUseCaseImpl) UpdateProject(project *models.ProjectCore) (err error) {
	return p.Gateway.UpdateProject(project)
}

func (p *ProjectUseCaseImpl) DeleteProject() {

}

func (p *ProjectUseCaseImpl) GetProjectById(projectId, userId string) (project *models.ProjectCore, err error) {
	project, getProjectErr := p.Gateway.GetProjectById(projectId, userId)
	if getProjectErr != nil {
		err = getProjectErr
		return
	}
	return
}
