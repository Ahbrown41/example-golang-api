package api

import (
	"api-reference-golang/api/entities"
	_ "api-reference-golang/docs" // docs is generated by Swag CLI, you have to import it
	"api-reference-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(conn *models.Connection) *gin.Engine {
	router := gin.Default()

	setupMonitoring(router)

	// Setup API Mappings
	entities.CreateRoutes(router, conn)

	// Setup Swagger Endpoint
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run
	router.Run()
	return router
}

// setupMonitoring - Sets up monitoring api
func setupMonitoring(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})

	// Setup Prometheus Metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	// Set slow time, default 5s
	m.SetSlowTime(10)
	// Set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(router)
}
