package http

import (
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, req interface{}) error {
	err := c.Bind(req)
	if err != nil {
		return err
	}

	return nil
}
