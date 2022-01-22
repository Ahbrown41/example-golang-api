package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ExecuteReq(t *testing.T, router *gin.Engine, verb string, url string, body interface{}) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	var bodyStr string
	if body != struct{}{} {
		bdy, err := json.Marshal(body)
		if err != nil {
			t.Errorf("Error marshalling object: %s", err)
		} else {
			bodyStr = string(bdy)
		}
	}
	req, err := http.NewRequest(verb, url, strings.NewReader(bodyStr))
	if err != nil {
		t.Errorf("Error creating new request: %s", err)
	}
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
