package delegate

import (
	"encoding/json"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
)

type CohortDelegateImpl struct {
	CohortUseCase cohorts.UseCase
	EdxUseCase    edx.UseCase
}

type CohortDelegateModule struct {
	fx.Out
	cohorts.Delegate
}

func SetupCohortDelegate(usecase cohorts.UseCase, edx edx.UseCase) CohortDelegateModule {
	return CohortDelegateModule{
		Delegate: &CohortDelegateImpl{
			usecase,
			edx,
		},
	}
}

func (p *CohortDelegateImpl) CreateCohort(cohort *models.CohortHTTP, courseId string) (newCohort models.CohortHTTP, err error) {
	cohortParams := models.CreateCohortHTTP{
		Message: map[string]interface{}{
			"name":            cohort.Name,
			"assignment_type": cohort.AssignmentType,
		},
	}

	body, err := p.EdxUseCase.CreateCohort(courseId, cohortParams.Message)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &cohort)
	if err != nil {
		log.Println(err)
		return
	}
	cohortCore := cohort.ToCore()
	newCohortCore, err := p.CohortUseCase.CreateCohort(cohortCore)
	if err != nil {
		log.Println(err)
		return
	}
	newCohort.FromCore(newCohortCore)
	return
}

func (p *CohortDelegateImpl) AddStudentToCohort(courseId, cohortId, studentId string) (err error) {
	_, err = p.EdxUseCase.AddStudentToCohort(courseId, cohortId, studentId)
	if err != nil {
		return
	}
	return
}
