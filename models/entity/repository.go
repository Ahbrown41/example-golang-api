package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reference-golang-svc/domainx"
)

//Repo : Database repo for DomainX Operations
type EntityRepository interface {
	// TASKS
	Fetch(opt FetchOptions) ([]string, error)
	GetByID(paperID int64) (string, error)
	GetBySourceID(paperID int64) (*domainx.Item, error)
	Create(o string) (int64, error)
	Update(o string) (string, error)
	Delete(paperID primitive.ObjectID) (bool, error)
}
