package project

import "gorm.io/gorm"

// Interface manages the project db mehtods
type Interface interface {
	Create(project *Project) error
	Delete(projectID uint) error
	GetByID(projectID uint) (*Project, error)
}

func New(db *gorm.DB) Interface {
	return &core{
		db: db,
	}
}

type core struct {
	db *gorm.DB
}

func (c core) Create(project *Project) error {
	return c.db.Create(project).Error
}

func (c core) Delete(projectID uint) error {
	return c.db.Delete(&Project{}, "id = ?", projectID).Error
}

func (c core) GetByID(projectID uint) (*Project, error) {
	prj := new(Project)

	query := c.db.
		First(&prj, "id = ?", projectID).
		Preload("Documents").
		Preload("Documents.Instructions")

	if err := query.Error; err != nil {
		return nil, ErrFailedToGetProject
	}

	if prj.ID != projectID {
		return nil, ErrProjectNotFound
	}

	return prj, nil
}
