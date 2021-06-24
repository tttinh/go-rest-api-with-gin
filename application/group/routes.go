package group

import "github.com/gin-gonic/gin"

func SetupHandler(router *gin.RouterGroup, endpoints Endpoints) {
	router.GET("/:id", endpoints.GetGroup)
}
