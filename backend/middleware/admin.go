package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/controller"
	"net/http"
)

func (m *middleware) Admin(c *gin.Context) {
	accessToken := m.service.Token.ExtractHeaderToken(c.Request)
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, controller.Error("authorization token not found in header"))
		c.Abort()
		return
	}
	key, err := m.service.Token.ExtractUserKey(accessToken, m.cts.AccessSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, controller.Error(err.Error()))
		c.Abort()
		return
	}

	// read users
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, controller.Error(err.Error()))
		c.Abort()
		return
	}
	if user.Admin != true {
		c.JSON(http.StatusUnauthorized, controller.Error("unauthorized access"))
		c.Abort()
		return
	}
	c.Next()
}
