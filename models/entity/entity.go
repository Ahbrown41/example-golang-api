package entity

import (
	"api-reference-golang/models"
	"api-reference-golang/models/notify"
	"encoding/json"
)

type service struct {
	repo Repository
	con  *models.Connection
	note notify.Notifier
}

type Repository interface {
	All(opt FetchOptions) (*[]Entity, error)
	Get(id uint) (*Entity, error)
	Create(entity *Entity) (*Entity, error)
	Delete(id uint) error
	Update(entity *Entity) (*Entity, error)
}

type FetchOptions struct {
	Limit int
	Page  int
}

func New(con *models.Connection, notify notify.Notifier) Repository {
	return service{con: con, note: notify}
}

// All : Returns all Entities
func (s service) All(opt FetchOptions) (*[]Entity, error) {
	var entities []Entity
	result := s.con.DB().Preload("Items").Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

// Get : Returns Entity by ID
func (s service) Get(id uint) (*Entity, error) {
	var entity Entity
	result := s.con.DB().Preload("Items").First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

// Create - Creates Entity
func (s service) Create(entity *Entity) (*Entity, error) {
	result := s.con.DB().Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	msg, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	s.note.Notify("createEntity", string(entity.ID), msg)
	return entity, nil
}

// Delete : Deletes Entity by ID
func (s service) Delete(id uint) error {
	var entity Entity
	result := s.con.DB().Preload("Items").First(&entity, id)
	if result.Error != nil {
		return result.Error
	}
	result = s.con.DB().Select("Items").Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	msg, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	s.note.Notify("deleteEntity", string(entity.ID), msg)
	return nil
}

// Update : Update Entity by ID
func (s service) Update(entity *Entity) (*Entity, error) {
	result := s.con.DB().Updates(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	msg, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	s.note.Notify("updateEntity", string(entity.ID), msg)
	return entity, nil
}
