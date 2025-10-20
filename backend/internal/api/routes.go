package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRoutes registra las rutas del API
func SetupRoutes(router *gin.Engine) {
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "backend ready",
		})
	})
}
