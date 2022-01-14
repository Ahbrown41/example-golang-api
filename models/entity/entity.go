package entity

import (
	"reference-golang-svc/models"
)

type EntityService struct {
	repo EntityRepository
	con  *models.Connection
}

type FetchOptions struct {
	Limit int
	Page  int
}

func New(con *models.Connection) *EntityService {
	return &EntityService{con: con}
}

// Fetch : Returns all Entities
func (s EntityService) All(opt FetchOptions) (*[]Entity, error) {
	var entities []Entity
	result := s.con.DB().Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

// Create - Creates Entity
func (s EntityService) Create(entity Entity) (Entity, error) {
	s.con.Validate(entity)
	result := s.con.DB().Create(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	return entity, nil
}
