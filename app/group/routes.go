package group

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, ctrl *Controller) {
	router := r.Group("/api/v1/group")
	router.GET("/:id", ctrl.GetGroup)
}
