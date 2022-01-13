package entity

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"reference-golang-svc/domainx"
)

type EntityService struct {
	repo EntityRepository
	db *gorm.DB
}

type FetchOptions struct {
	Limit int
	Page int
}

func New(db *gorm.DB) *EntityService {
	return &EntityService{db: db}
}

// Fetch : Returns all Papers
func (s EntityService) All(opt FetchOptions) ([]*Entity, error) {
	var entities []Entity
	result := s.db.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}


	query := bson.M{}
	opts := options.FindOptions{Limit: &opt.Limit, Skip: &opt.Page}
	cur, err := col().Find(context.TODO(), query, &opts)
	if result.Error != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var papers []*domainx.Item
	err = cur.All(context.TODO(), &papers)
	if err != nil {
		return nil, err
	}
	return papers, nil
}

// GetByID : Returns a paper by ID
func (s EntityService) GetByID(paperID primitive.ObjectID) (*domainx.Item, error) {
	resultDoc := col().FindOne(context.TODO(), bson.M{"_id": paperID})
	paper := &domainx.Item{}
	err := resultDoc.Decode(paper)
	if err != nil {
		return nil, err
	}
	return paper, nil
}

// Create - Creates Taskpaper
func (s EntityService) Create(paper string) (primitive.ObjectID, error) {
	paperObj, _ := domainx.Unmarshal([]byte(paper))
	paperObj.ID = primitive.NewObjectID()
	paperObj.Parent = &domainx.Item{}

	res, err := col().InsertOne(context.TODO(), *paperObj)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// Update - Updates paper
func (s EntityService) Update(paper string) (int64, error) {
	paperObj, _ := domainx.Unmarshal([]byte(paper))
	res, err := col().UpdateOne(
		context.Background(),
		bson.M{"_id": paperObj.ID},
		bson.M{"$set": paper},
	)
	if err != nil {
		return 0, err
	}
	if (res.ModifiedCount + res.MatchedCount) > 0 {
		return res.ModifiedCount, nil
	}
	return 0, errors.New("Invalid Object ID")
}

// Delete - Deletes paper
func (s EntityService) Delete(paperID primitive.ObjectID) (int64, error) {
	dbClient.


	res, err := col().DeleteOne(context.Background(), bson.M{"_id": paperID})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
