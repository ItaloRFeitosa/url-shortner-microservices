package link_management

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	c := NewClient()
	h := NewHandler(c)
	routes := r.Group("/links")
	routes.POST("/", h.CreateLink)
	// routes.GET("/:id", svc.FindOne)
}
