package entities

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io/ioutil"
	"net/http"
)

// AllEntities godoc
// @Summary Gets all Entities
// @Description Gets all Entities in DB
// @Tags Entities
// @Accept  json
// @Produce  json
// @Success 200 {array} Entities.Item
// @Router /Entities/ [get]
func AllEntities(c *gin.Context) {
	opts := Entities.FetchOptions{
		Page:  com.StrTo(c.DefaultQuery("page", "0")).MustInt64(),
		Limit: com.StrTo(c.DefaultQuery("limit", "100")).MustInt64(),
	}
	Entities, _ := Entities.Entitieservice{}.Fetch(opts)
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(Entities)
}

// GetEntityByID godoc
// @Summary Gets a given Entity JSON
// @Description Gets a given Entity in a Entity BSON format
// @Tags Entities
// @Accept  json
// @Produce  json
// @Success 200 {object} Entities.Item
// @Router /Entities/:pid [get]
func GetEntityByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	Entities, _ := Entities.Entitieservice{}.GetByID(getIDfromHex(c.Param("pid")))
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
	c.Header("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(c.Request.Body)
	pid, _ := Entities.Entitieservice{}.Create(string(body))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": pid.Hex(),
	})
}

// UpdateEntity godoc
// @Summary Updates a Entities document
// @Description Updates a Entity given an Entity ID and payload
// @Tags Entities
// @Accept  json
// @Produce  json
// @Param pid path string true "Entity to update"
// @Param Entities body Entities.Item true "Updated Entity"
// @Success 200 {string} string "Updated Entity"
// @Router /Entities/:pid [put]
func UpdateEntity(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(c.Request.Body)
	cnt, _ := Entities.Entitieservice{}.Update(string(body))
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"ModifiedCount": cnt,
	})
}

// DeleteEntity godoc
// @Summary Deletes a Entities Entities
// @Description Deletes an Entities given an Entities ID
// @Tags Entities
// @Accept  json
// @Produce  json
// @Param pid path string true "Task to delete"
// @Success 200 "Updated Task"
// @Router /Entities/:pid [delete]
func DeleteEntity(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	cnt, _ := Entities.Entitieservice{}.Delete(getIDfromHex(c.Param("pid")))
	c.JSON(http.StatusOK, map[string]interface{}{
		"DeleteCount": cnt,
	})
}
