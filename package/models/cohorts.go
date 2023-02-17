package models

import "gorm.io/gorm"

type CreateCohortHTTP struct {
	Message map[string]interface{} `json:"message"`
}

type CohortDB struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex"`
	UserCount       int
	AssignmentType  string
	UserPartitionID int `gorm:"default:null"`
	GroupID         int `gorm:"default:null"`
}

type CohortCore struct {
	ID              int
	Name            string
	UserCount       int
	AssignmentType  string
	UserPartitionID int
	GroupID         int
}

func (ht *CohortHTTP) ToCore() *CohortCore {
	return &CohortCore{
		ID:              ht.ID,
		Name:            ht.Name,
		UserCount:       ht.UserCount,
		AssignmentType:  ht.AssignmentType,
		UserPartitionID: ht.UserPartitionID,
		GroupID:         ht.GroupID,
	}
}

func (ht *CohortHTTP) FromCore(core *CohortCore) {
	ht.ID = core.ID
	ht.Name = core.Name
	ht.UserCount = core.UserCount
	ht.AssignmentType = core.AssignmentType
	ht.UserPartitionID = core.UserPartitionID
	ht.GroupID = core.GroupID
}

func (em *CohortDB) ToCore() *CohortCore {
	return &CohortCore{
		ID:              int(em.ID),
		Name:            em.Name,
		UserCount:       em.UserCount,
		AssignmentType:  em.AssignmentType,
		UserPartitionID: em.UserPartitionID,
		GroupID:         em.GroupID,
	}
}

func (em *CohortDB) FromCore(core *CohortCore) {
	em.ID = uint(core.ID)
	em.Name = core.Name
	em.UserCount = core.UserCount
	em.AssignmentType = core.AssignmentType
	em.UserPartitionID = core.UserPartitionID
	em.GroupID = core.GroupID
}
