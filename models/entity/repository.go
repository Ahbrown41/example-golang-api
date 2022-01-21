package entity

// EntityRepository : Database repo for Domain Operations
type EntityRepository interface {
	All(opt FetchOptions) ([]string, error)
	Create(entity Entity) (Entity, error)
}
