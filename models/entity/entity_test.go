package entity

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"reference-golang-svc/domainx"
	"reference-golang-svc/repos"
	"reflect"
	"testing"
)

var svc PaperService
var paper1Obj *domainx.Item

func TestMain(t *testing.M) {
	repos.InitDB("mongodb://localhost:27017", "tasks-test")
	createPaper()
	paper1Obj, _ = svc.GetByID(_id)
	code := t.Run()
	defer repos.DisconnectDB()
	clearTable()
	os.Exit(code)
}

func createPaper() {
	svc := PaperService{}
	id, err := svc.Create(paper1)
	if err != nil {
		log.Fatal(err)
	}
	_id = id
	fmt.Printf("Created object with id: %s\n", _id)
}

func TestFind(t *testing.T) {
	paper, err := svc.GetByID(_id)
	assert.Equal(t, err, nil, "Error not nil")
	assert.NotEqual(t, nil, paper, "They should not be equal")
	assert.Equal(t, true, reflect.DeepEqual(paper, paper1Obj), fmt.Sprintf("Paper: %v != %v", paper, paper1Obj))
}

func TestFetch(t *testing.T) {
	options := FetchOptions{
		Limit: 100,
		Page:  0,
	}
	papers, err := svc.Fetch(options)
	assert.Equal(t, err, nil, "Error not nil")
	assert.NotEqual(t, nil, papers, "They should not be equal")
	assert.True(t, len(papers) == 1, "They should be greater than zero but was: %d", len(papers))
	for _, paper := range papers {
		assert.Equal(t, true, reflect.DeepEqual(paper, paper1Obj), fmt.Sprintf("Paper: %v != %v", paper, paper1Obj))
	}
}

func TestDelete(t *testing.T) {
	cnt, err := svc.Delete(_id)
	assert.Equal(t, err, nil, "Error not nil")
	assert.NotEqual(t, cnt, 0, "They should not be equal")
}

func clearTable() {
	repos.GetCollection("entities").DeleteMany(context.TODO(), bson.M{})
}
