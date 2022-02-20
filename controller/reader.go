package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service/suitetalk"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

// ReadImageText read image from file and return text
func (o *controller) ReadImageText(c *gin.Context) {
	f, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	text, err := o.Tesseract.ReadImage(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(text))
}

// PostImage post image to NetSuite
func (o *controller) PostImage(c *gin.Context) {
	template := c.PostForm("template")
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnauthorized, Error(err.Error()))
		return
	}
	open, err := f.Open()
	if err != nil {
		c.JSON(http.StatusUnauthorized, Error(err.Error()))
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, open); err != nil {
		c.JSON(http.StatusUnauthorized, Error(err.Error()))
		return
	}
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)
	config, err := o.Storage.ReadConfig()
	if err != nil {
		c.JSON(http.StatusUnauthorized, Error(err.Error()))
		return
	}
	if config.WsUrl == "" || config.Realm == "" || config.Version == "" {
		c.JSON(http.StatusForbidden, Error("company configuration values"))
		return
	}
	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		c.JSON(http.StatusForbidden, Error("consumer configuration missing"))
		return
	}
	if config.TokenID == "" || config.TokenSecret == "" {
		c.JSON(http.StatusForbidden, Error("token configuration missing"))
		return
	}
	file := suitetalk.NewFile(template, f.Filename, buf)
	settings := suitetalk.NewSettings(config)
	upload, err := suitetalk.Upload(settings, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(upload))
}
