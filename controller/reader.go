package controller

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/suitetalk"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type ImagePost struct {
	Template string `form:"template"`
}

// ReadImageText read image from file and return text
func (o *controller) ReadImageText(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	text, err := o.Tesseract.ReadImage(b)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(text)
}

// PostImage post image to NetSuite
func (o *controller) PostImage(c *fiber.Ctx) error {
	var imagePost ImagePost
	err := c.BodyParser(imagePost)
	if err != nil {
		return err
	}
	template := imagePost.Template
	if template == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No template sent")
	}
	f, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	open, err := f.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, open); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)
	config, err := o.Storage.ReadConfig()
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if config.WsUrl == "" || config.Realm == "" || config.Version == "" {
		return fiber.NewError(fiber.StatusForbidden, "Company settings missing")
	}
	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		return fiber.NewError(fiber.StatusForbidden, "Consumer keys missing")
	}
	if config.TokenID == "" || config.TokenSecret == "" {
		return fiber.NewError(fiber.StatusForbidden, "Token keys missing")
	}
	file := suitetalk.NewFile(template, f.Filename, buf)
	settings := suitetalk.NewSettings(config)
	upload, err := suitetalk.Upload(settings, file)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(upload)
}
