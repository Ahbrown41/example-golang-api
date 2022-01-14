package entities

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io/ioutil"
	"net/http"
	"reference-golang-svc/endpoints"
	"reference-golang-svc/models/entity"
)

// AllEntities godoc
// @Summary Gets all Entities
// @Description Gets all Entities in DB
// @Tags Entities
// @Accept  json
// @Produce  json
// @Success 200 {array} Entity.Item
// @Router /entities/ [get]
func AllEntities(c *gin.Context) {
	opts := entity.FetchOptions{
		Page:  com.StrTo(c.DefaultQuery("page", "0")).MustInt(),
		Limit: com.StrTo(c.DefaultQuery("limit", "100")).MustInt(),
	}
	Entities, _ := entity.EntityService{}.All(opts)
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(Entities)
}

// CreateEntity godoc
// @Summary Creates Entity document based on RAW Entities
// @Description Creates a Entities document
// @Tags Entities
// @Accept  json
// @Produce  json
// @Param Entities body string true "New Entity"
// @Success 200 {string} string "New Entity ID(Hex)"
// @Router /Entities/ [post]
func CreateEntity(c *gin.Context) {
	var ety entity.Entity
	c.BindJSON(ety)
	entity, err := entity.EntityService{}.Create(ety)
	if err != nil {
		endpoints.RespondJSON(c, 404, nil)
	} else {
		endpoints.RespondJSON(c, 200, entity)
	}

	c.Header("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(c.Request.Body)
	pid, _ := Entities.Entitieservice{}.Create(string(body))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": pid.Hex(),
	})
}
