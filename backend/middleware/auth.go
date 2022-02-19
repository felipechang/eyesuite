package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/controller"
	"gitlab.com/hardcake/eyesuite/service/token"
	"net/http"
)

func (m *middleware) Auth(c *gin.Context) {

	// find and evaluate access token
	key, err := m.evaluateAccessKey(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, controller.Error(err.Error()))
		c.Abort()
		return
	}
	if key != "" {
		c.Set("user_key", key)
		c.Next()
		return
	}

	// find and evaluate refresh token
	key, err = m.evaluateRefreshKey(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Error(err.Error()))
		c.Abort()
		return
	}
	if key != "" {
		c.Set("user_key", key)
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, controller.Error("could not authenticate tokens"))
	c.Abort()
}

// evaluateAccessKey evaluate that access token is correct
func (m *middleware) evaluateAccessKey(c *gin.Context) (string, error) {
	headerToken := m.service.Token.ExtractHeaderToken(c.Request)
	if headerToken == "" {
		return "", nil
	}
	key, err := m.service.Token.ExtractUserKey(headerToken, m.cts.AccessSecret)
	if err != nil {
		return "", err
	}
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	if user.AccessToken != headerToken {
		return "", errors.New("invalid token for user")
	}
	return key, nil
}

// evaluateRefreshKey evaluate that refresh token is correct
func (m *middleware) evaluateRefreshKey(c *gin.Context) (string, error) {
	req := &token.Pair{}
	err := c.ShouldBindJSON(&req)
	if err != nil && err.Error() != "EOF" {
		return "", errors.New("could not parse body token")
	}
	if req.RefreshToken == "" {
		return "", errors.New("refresh token not found")
	}

	key, err := m.service.Token.ExtractUserKey(req.RefreshToken, m.cts.RefreshSecret)
	if err != nil {
		return "", err
	}
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	if user.RefreshToken != req.RefreshToken {
		return "", errors.New("invalid token for user")
	}
	return key, nil
}
