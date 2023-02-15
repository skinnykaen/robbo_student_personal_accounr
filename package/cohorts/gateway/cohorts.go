package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type CohortsGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type CohortsGatewayModule struct {
	fx.Out
	cohorts.Gateway
}

func SetupCohortsGateway(postgresClient db_client.PostgresClient) CohortsGatewayModule {
	return CohortsGatewayModule{
		Gateway: &CohortsGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *CohortsGatewayImpl) CreateCohort(cohortCore *models.CohortCore) (newCohort *models.CohortCore, err error) {
	cohortDb := models.CohortDB{}
	cohortDb.FromCore(cohortCore)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&cohortDb).Error
		return
	})

	newCohort = cohortDb.ToCore()
	return
}
