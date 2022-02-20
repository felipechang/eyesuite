package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"net/http"
)

// ReadPlugins read plugin list
func (o *controller) ReadPlugins(c *gin.Context) {
	rp, err := o.Storage.ReadPlugins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(rp))
}

// UpsertPlugins update or create plugin list
func (o *controller) UpsertPlugins(c *gin.Context) {
	var plugins []storage.Plugin
	if err := c.ShouldBindJSON(&plugins); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	up, err := o.Storage.UpsertPlugins(plugins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(up))
}
