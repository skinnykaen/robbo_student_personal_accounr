package delegate

import (
	"encoding/json"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"strconv"
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

func (p *CohortDelegateImpl) CreateCohort(cohort *models.CohortHTTP, createCohort *models.CreateCohortHTTP, courseId string) (id string, err error) {
	body, err := p.EdxUseCase.CreateCohort(courseId, createCohort.Message)
	if err != nil {
		return "", cohorts.ErrBadRequest
	}
	err = json.Unmarshal(body, cohort)
	if err != nil {
		return "", cohorts.ErrInternalServerLevel
	}
	id = strconv.FormatUint(uint64(cohort.ID), 10)
	return
}

func (p *CohortDelegateImpl) AddStudent(username, courseId string, cohortId int) (err error) {
	_, err = p.EdxUseCase.AddStudent(username, courseId, cohortId)
	if err != nil {
		return cohorts.ErrBadRequest
	}
	return
}
