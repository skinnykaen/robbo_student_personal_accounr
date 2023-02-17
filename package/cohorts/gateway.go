package cohorts

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateCohort(cohortCore *models.CohortCore) (newCohort *models.CohortCore, err error)
}
