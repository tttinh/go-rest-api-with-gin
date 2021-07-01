package group

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, service Service) {
	router := r.Group("/api/v1/group")

	ctrl := NewController(service)
	router.GET("/:id", ctrl.GetGroup)
	router.POST("/", ctrl.CreateGroup)
}
