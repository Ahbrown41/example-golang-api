package entity

import (
	"api-reference-golang/models"
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

// All : Returns all Entities
func (s EntityService) All(opt FetchOptions) (*[]Entity, error) {
	var entities []Entity
	result := s.con.DB().Preload("Items").Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

// Get : Returns Entity by ID
func (s EntityService) Get(id uint) (*Entity, error) {
	var entity Entity
	result := s.con.DB().Preload("Items").First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

// Create - Creates Entity
func (s EntityService) Create(entity Entity) (Entity, error) {
	result := s.con.DB().Create(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	return entity, nil
}

// Delete : Deletes Entity by ID
func (s EntityService) Delete(id uint) (*Entity, error) {
	var entity Entity
	result := s.con.DB().Delete(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

// Update : Update Entity by ID
func (s EntityService) Update(entity Entity) (*Entity, error) {
	result := s.con.DB().Updates(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}
