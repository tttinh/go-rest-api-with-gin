package group

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, service Service) {
	router := r.Group("/api/v1/groups")

	ctrl := NewController(service)
	router.GET("/:id", ctrl.GetGroup)
	router.PUT("/:id", ctrl.UpdateGroup)
	router.DELETE("/:id", ctrl.DeleteGroup)
	router.POST("/", ctrl.CreateGroup)
}
