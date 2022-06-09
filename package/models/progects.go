package models

import (
	"gorm.io/gorm"
	"strconv"
)

type ProjectCore struct {
	ID     string
	Name   string
	Author string
	Json   string
}

type ProjectPageCore struct {
	ProjectsCore []*ProjectCore
	Description  string
	Preview      string
	LinkScratch  string
}

type ProjectDB struct {
	gorm.Model

	Name   string `gorm:"not null;size:256"`
	Author string `gorm:"not null;size:256"`
	Json   string `gorm:"not null;size:65535"`
}

type ProjectPageDB struct {
	gorm.Model

	PP          ProjectDB `gorm:"foreignKey:PPId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string    `gorm:"size:256;not null"`
	Preview     string    `gorm:"size:256;not null"`
	LinkScratch string    `gorm:"size:256;not null"`
}

type ProjectHTTP struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Json   string    `json:"json"`
}

type ProjectPageHTTP struct {
	ProjectsHTTP []*ProjectHTTP `json:"projects"`
	Description  string         `json:"description"`
	Preview      string         `json:"preview"`
	LinkScratch  string         `json:"link"`
}

func (em *ProjectDB) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:     strconv.FormatUint(uint64(em.ID), 10),
		Name:   em.Name,
		Author: em.Author,
		Json:   em.Json,
	}
}

func (em *ProjectDB) FromCore(project *ProjectCore) {
	id, _ := strconv.ParseUint(project.ID, 10, 64)
	em.ID = uint(id)
	em.Name = project.Name
	em.Author = project.Author
	em.Json = project.Json
}

func (em *ProjectPageDB) ToCore(projects []*ProjectDB) ProjectPageCore {
	var coreProjects []*ProjectCore
	for _, projectDB := range projects {
		coreProjects = append(coreProjects, projectDB.ToCore())
	}
	return ProjectPageCore{
		Description:  em.Description,
		Preview:      em.Preview,
		LinkScratch:  em.LinkScratch,
		ProjectsCore: coreProjects,
	}
}

func (em *ProjectPageDB) FromCore(pp *ProjectPageCore) {
	em.Description = pp.Description
	em.Preview = pp.Preview
	em.LinkScratch = pp.LinkScratch
}

func (ht *ProjectHTTP) ToCore() *ProjectCore {
	return &ProjectCore{
		ID:     ht.ID,
		Name:   ht.Name,
		Author: ht.Author,
		Json:   ht.Json,
	}
}

func (ht *ProjectHTTP) FromCore(project *ProjectCore) {
	ht.ID = project.ID
	ht.Name = project.Name
	ht.Author = project.Author
	ht.Json = project.Json
}

func (ht *ProjectPageHTTP) ToCore(projects []*ProjectDB) ProjectPageCore {
	var coreProjects []*ProjectCore
	for _, projectDB := range projects {
		coreProjects = append(coreProjects, projectDB.ToCore())
	}
	return ProjectPageCore{
		Description:  ht.Description,
		Preview:      ht.Preview,
		LinkScratch:  ht.LinkScratch,
		ProjectsCore: coreProjects,
	}
}

func (ht *ProjectPageHTTP) FromCore(pp *ProjectPageCore) {
	ht.Description = pp.Description
	ht.Preview = pp.Preview
	ht.LinkScratch = pp.LinkScratch
}
