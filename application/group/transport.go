package group

import (
	"github.com/gin-gonic/gin"
)

func MakeHandler(router *gin.RouterGroup, groupService Service) {

	router.GET("/:id", makeGetGroupEndpoint(groupService))

}
