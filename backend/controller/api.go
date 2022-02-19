package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *controller) ApiHome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to EyeSuite")
}
