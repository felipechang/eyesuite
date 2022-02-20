package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service"
)

type Controller interface {
	ApiHome(c *gin.Context)

	UpsertConfig(c *gin.Context)
	ReadConfig(c *gin.Context)

	Login(c *gin.Context)
	Refresh(c *gin.Context)
	Logout(c *gin.Context)

	ReadProfiles(c *gin.Context)
	UpsertProfiles(c *gin.Context)

	ReadImageText(c *gin.Context)
	PostImage(c *gin.Context)

	ReadUsers(c *gin.Context)
	UpsertUsers(c *gin.Context)

	ReadPlugins(c *gin.Context)
	UpsertPlugins(c *gin.Context)
}

type controller struct {
	*service.Service
}

func NewController(s *service.Service) Controller {
	return &controller{s}
}
