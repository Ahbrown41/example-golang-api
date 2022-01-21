package entities

import (
	"api-reference-golang/api/utils"
	"api-reference-golang/models/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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
	entities, _ := entity.EntityService{}.All(opts)
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(entities)
}

// GetEntityByID godoc
// @Summary Gets and entity by ID
// @Description Gets and entity by ID
// @Tags Entities
// @Accept  json
// @Produce  json
// @Success 200 Entity.Item
// @Router /entities/:id [get]
func GetEntityByID(c *gin.Context) {
	entity, _ := entity.EntityService{}.Get(uint(c.GetInt("id")))
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(entity)
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
	err := c.BindJSON(&ety)
	if err != nil {
		utils.RespondJSON(c, 404, nil)
	}
	entity, err := entity.EntityService{}.Create(ety)
	c.Header("Content-Type", "application/json")
	if err != nil {
		utils.RespondJSON(c, 404, nil)
	} else {
		utils.RespondJSON(c, 200, entity)
	}
}

func CreateRoutes(router *gin.Engine) {
	pv1 := router.Group("/api/v1/entities")
	{
		pv1.POST("/", CreateEntity)
		pv1.GET("/", AllEntities)
		pv1.GET("/:pid", GetEntityByID)
		//pv1.PUT("/:pid", domain.UpdateEntity)
		//pv1.DELETE("/:pid", domain.DeleteEntity)
	}
}
