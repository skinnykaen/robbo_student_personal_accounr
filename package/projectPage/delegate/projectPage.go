package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"go.uber.org/fx"
)

type ProjectPageDelegateImpl struct {
	projectPage.UseCase
}

type ProjectPageDelegateModule struct {
	fx.Out
	projectPage.Delegate
}

func SetupProjectPageDelegate(usecase projectPage.UseCase) ProjectPageDelegateModule {
	return ProjectPageDelegateModule{
		Delegate: &ProjectPageDelegateImpl{
			UseCase: usecase,
		},
	}
}

func (p *ProjectPageDelegateImpl) CreateProjectPage() {

}

func (p *ProjectPageDelegateImpl) DeleteProjectPage(projectID string) (err error) {
	return p.UseCase.DeleteProjectPage(projectID)
}

func (p *ProjectPageDelegateImpl) UpdateProjectPage(projectPage *models.ProjectPageHTTP) (err error) {
	projectPageCore := projectPage.ToCore()
	return p.UseCase.UpdateProjectPage(projectPageCore)
}

func (p *ProjectPageDelegateImpl) GetProjectPageByID(projectID string) (projectPage *models.ProjectPageHTTP, err error) {
	projectPageCore, err := p.UseCase.GetProjectPageByID(projectID)
	projectPage.FromCore(projectPageCore)
	return
}
