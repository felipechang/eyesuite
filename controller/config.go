package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"net/http"
)

// ReadConfig read configuration
func (o *controller) ReadConfig(c *gin.Context) {
	rc, err := o.Storage.ReadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(rc))
}

// UpsertConfig update or create a configuration
func (o *controller) UpsertConfig(c *gin.Context) {
	config := &storage.Config{}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	uc, err := o.Storage.UpsertConfig(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(uc))
}
