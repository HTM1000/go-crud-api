package router

import (
	"github.com/HTM1000/go-oportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// Show opening
		v1.GET("/opening", handler.ShowOpeningHandler)

		// Create opening
		v1.POST("/opening", handler.CreateOpeningHandler)

		// Delete opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// Update opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// List openings
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
