package entities

import (
	"api-reference-golang/models"
	"api-reference-golang/models/entity"
)

type service struct {
}

type Service interface {
	Create(entity *entity.Entity) (*entity.Entity, error)
}

func New(con *models.Connection) Service {
	return service{}
}

// All : Returns all Entities
func (s service) Create(entity *entity.Entity) (*entity.Entity, error) {
	return nil, nil
}
