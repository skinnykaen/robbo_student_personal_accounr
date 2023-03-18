package robboGroup

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateRobboGroup(robboGroupCore *models.RobboGroupCore) (newRobboGroup *models.RobboGroupCore, err error)
	UpdateRobboGroup(robboGroupCore *models.RobboGroupCore) (robboGroupUpdated *models.RobboGroupCore, err error)
	DeleteRobboGroup(robboGroupId string) (err error)
	GetAllRobboGroups(page, pageSize int) (robboGroupsCore []*models.RobboGroupCore, countRows int64, err error)
	GetRobboGroupsByRobboUnitId(robboUnitId string) (robboGroupsCore []*models.RobboGroupCore, err error)
	GetRobboGroupsByRobboUnitsIds(robboUnitsId []string, page, pageSize int) (robboGroupsCore []*models.RobboGroupCore, countRows int64, err error)
	GetRobboGroupById(robboGroupId string) (robboGroupCore *models.RobboGroupCore, err error)
	SetTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteTeacherForRobboGroup(relation *models.TeachersRobboGroupsCore) (err error)
	DeleteRelationByRobboGroupId(robboGroupId string) (err error)
	DeleteRelationByTeacherId(teacherId string) (err error)
	GetRelationByRobboGroupId(robboGroupId string) (relations []*models.TeachersRobboGroupsCore, err error)
	GetRelationByTeacherId(teacherId string, page, pageSize int) (relations []*models.TeachersRobboGroupsCore, countRows int64, err error)
	SearchRobboGroupsByTitle(title string, page, pageSize int) (robboGroupsCore []*models.RobboGroupCore, countRows int64, err error)
}
