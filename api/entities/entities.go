package entities

import (
	"api-reference-golang/api/utils"
	"api-reference-golang/models"
	"api-reference-golang/models/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var (
	svc *entity.EntityService
)

// AllEntities godoc
// @Summary      Gets all Entities
// @Description  Gets all Entities in DB
// @Tags         Entities
// @Accept       json
// @Produce      json
// @Success      200  {array} entity.Entity
// @Router       /api/v1/entities/ [get]
func AllEntities(c *gin.Context) {
	opts := entity.FetchOptions{
		Page:  com.StrTo(c.DefaultQuery("page", "0")).MustInt(),
		Limit: com.StrTo(c.DefaultQuery("limit", "100")).MustInt(),
	}
	entities, _ := svc.All(opts)
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(entities)
}

// GetEntityByID godoc
// @Summary      Gets an entity by ID
// @Description  Gets an entity by ID
// @Tags         Entities
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Entity ID to load"
// @Success      200  {object} entity.Entity
// @Failure      404  {object}  utils.ResponseData  "Can not find ID"
// @Router       /api/v1/entities/{id} [get]
func GetEntityByID(c *gin.Context) {
	entity, _ := svc.Get(uint(utils.ParamInt64(c, "id")))
	c.Header("Content-Type", "application/json")
	if entity != nil {
		c.JSON(200, entity)
	} else {
		utils.RespondJSON(c, 404, entity)
	}
}

// CreateEntity godoc
// @Summary      Creates new Entity
// @Description  Creates new Entity
// @Tags         Entities
// @Accept       json
// @Produce      json
// @Param        entity	body	entity.Entity  true  "Entity to update"
// @Success      200 {object}  entity.Entity  "New Entity"
// @Router       /api/v1/entities/ [post]
func CreateEntity(c *gin.Context) {
	var ety entity.Entity
	err := c.BindJSON(&ety)
	if err != nil {
		utils.RespondJSON(c, 400, err.Error())
		return
	}
	entity, err := svc.Create(&ety)
	c.Header("Content-Type", "application/json")
	if err != nil {
		utils.RespondJSON(c, 400, err.Error())
	} else {
		c.JSON(200, entity)
	}
	return
}

// UpdateEntity godoc
// @Summary      Updates an entity
// @Description  Updates an entity
// @Tags         Entities
// @Accept       json
// @Produce      json
// @Param        entity	body	entity.Entity  true  "Entity to update"
// @Success      200 {object} entity.Entity
// @Router       /api/v1/entities/{id} [put]
func UpdateEntity(c *gin.Context) {
	var ety entity.Entity
	err := c.BindJSON(&ety)
	if err != nil {
		utils.RespondJSON(c, 404, nil)
	}
	entity, err := svc.Update(&ety)
	c.Header("Content-Type", "application/json")
	if err != nil {
		utils.RespondJSON(c, 404, nil)
	} else {
		c.JSON(200, entity)
	}
}

// DeleteEntity godoc
// @Summary      Deletes an entity
// @Description  Deletes an entity
// @Tags         Entities
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Entity ID to delete"
// @Success      200  {object}  entity.Entity
// @Failure      404  {object}  utils.ResponseData  "Can not find ID"
// @Router       /api/v1/entities/{id} [delete]
func DeleteEntity(c *gin.Context) {
	err := svc.Delete(uint(utils.ParamInt64(c, "id")))
	if err != nil {
		utils.RespondJSON(c, 404, nil)
	} else {
		utils.RespondJSON(c, 200, nil)
	}
}

func CreateRoutes(router *gin.Engine, conn *models.Connection) {
	svc = entity.New(conn)
	pv1 := router.Group("/api/v1/entities")
	{
		pv1.POST("/", CreateEntity)
		pv1.GET("/", AllEntities)
		pv1.GET("/:id", GetEntityByID)
		pv1.PUT("/:id", UpdateEntity)
		pv1.DELETE("/:id", DeleteEntity)
	}
}
