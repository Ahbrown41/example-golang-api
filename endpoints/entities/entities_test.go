package entities

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reference-golang-svc/domainx"
	"reference-golang-svc/repos"
	"strings"
	"testing"
)

const baseURL = "/api/v1/entities"

var router *gin.Engine

const paper1 = `My Project:
			- Task A
				- Task B
			- Task C
				Some note for Task C
		Project #2:
			- Project 2 Task 1 #fish
		`

func setupRouter() *gin.Engine {
	router := gin.Default()
	pv1 := router.Group("/api/v1/entities")
	{
		pv1.POST("/", CreatePaperEndPoint)
		pv1.GET("/", AllPapersEndPoint)
		pv1.GET("/:pid", GetPaperByIDEndpoint)
		pv1.GET("/:pid/raw", GetPaperRawByIDEndpoint)
		pv1.PUT("/:pid", UpdatePaperEndPoint)
		pv1.DELETE("/:pid", DeletePaperEndPoint)
	}
	return router
}

func TestMain(m *testing.M) {

	repos.InitDB("mongodb://localhost:27017", "recontrac-test")
	router = setupRouter()
	code := m.Run()
	defer repos.DisconnectDB()
	os.Exit(code)
}

func createPaper(t *testing.T) {
	url := fmt.Sprintf("%s/", baseURL)
	rr := executeReq(t, "POST", url, strings.NewReader(paper1), CreatePaperEndPoint)
	assert.Equal(t, http.StatusOK, rr.Code, "Return code not 200")
	var result map[string]interface{}
	json.Unmarshal([]byte(rr.Body.String()), &result)

	assert.NotNil(t, result, "Papers not nil")
	assert.NotNil(t, result["id"], "Paper ID not valid")

	paperID = getIDfromHex(result["id"].(string))
}

func TestGetPapersEndpoint(t *testing.T) {
	url := fmt.Sprintf("%s/", baseURL)

	createPaper(t)

	rr := executeReq(t, "GET", url, nil, AllPapersEndPoint)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var results []map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &results)
	assert.NotNil(t, results, "Papers is nil")
	assert.Greater(t, len(results), 0, "Error body does not equal 1")
}

func TestGetPaperEndpoint(t *testing.T) {
	url := fmt.Sprintf("%s/%s", baseURL, paperID.Hex())
	rr := executeReq(t, "GET", url, nil, GetPaperByIDEndpoint)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var paper domainx.Item
	_ = json.NewDecoder(rr.Body).Decode(&paper)
	assert.NotNil(t, paper, "Paper is nil")
	assert.Equal(t, paper.ID.Hex(), paperID.Hex(), "Paper ID do not match")
}

func TestGetPaperRawEndpoint(t *testing.T) {
	url := fmt.Sprintf("%s/%s/raw", baseURL, paperID.Hex())
	rr := executeReq(t, "GET", url, nil, GetPaperByIDEndpoint)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	assert.Equal(t, paper1, rr.Body.String(), "Paper ID do not match")
}

func TestDeletePaperEndpoint(t *testing.T) {
	url := fmt.Sprintf("%s/%s", baseURL, paperID.Hex())
	rr := executeReq(t, "DELETE", url, nil, DeletePaperEndPoint)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var results map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &results)
	assert.Equal(t, float64(1), results["DeleteCount"], "Error body does not equal 1")
}

func executeReq(t *testing.T, verb string, url string, body io.Reader, endpoint func(ctx *gin.Context)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(verb, url, body)
	router.ServeHTTP(rr, req)

	checkStatus(t, rr)
	return rr
}

func checkStatus(t *testing.T, rr *httptest.ResponseRecorder) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
