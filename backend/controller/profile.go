package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"net/http"
)

// ReadProfiles read profile list
func (o *controller) ReadProfiles(c *gin.Context) {
	rp, err := o.Storage.ReadProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(rp))
}

// UpsertProfiles update or create profile list
func (o *controller) UpsertProfiles(c *gin.Context) {
	var profiles []storage.Profile
	if err := c.ShouldBindJSON(&profiles); err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	up, err := o.Storage.UpsertProfiles(profiles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(up))
}
