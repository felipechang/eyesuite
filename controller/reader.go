package controller

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/makiuchi-d/gozxing/qrcode"

	"gitlab.com/hardcake/eyesuite/service/suitetalk"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
)

// getFileImage read an image and return its contents
func (o *controller) getFileImage(c *fiber.Ctx) ([]byte, error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PostImage post image to NetSuite
func (o *controller) PostImage(c *fiber.Ctx) error {
	template := c.FormValue("template")
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

// ReadImageText read a text image from file and return text
func (o *controller) ReadImageText(c *fiber.Ctx) error {
	b, err := o.getFileImage(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	text, err := o.Tesseract.ReadImage(b)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(text)
}

// ReadImageBarcode read a barcode image from file and return text
func (o *controller) ReadImageBarcode(c *fiber.Ctx) error {
	v := c.FormValue("barcode-type")
	if v == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "barcode-type not set")
	}
	if v != "upca" && v != "upce" && v != "ean8" && v != "ean13" {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("invalid barcode-type %s", v))
	}
	b, err := o.getFileImage(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	var codeReader gozxing.Reader
	if v == "upca" {
		codeReader = oned.NewUPCAReader()
	}
	if v == "upce" {
		codeReader = oned.NewUPCEReader()
	}
	if v == "ean8" {
		codeReader = oned.NewEAN8Reader()
	}
	if v == "ean13" {
		codeReader = oned.NewEAN13Reader()
	}
	result, err := codeReader.Decode(bmp, nil)
	if err != nil {
		return err
	}
	return c.JSON(result.GetText())
}

// ReadImageQr read a QR image from file and return text
func (o *controller) ReadImageQr(c *fiber.Ctx) error {
	b, err := o.getFileImage(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	codeReader := qrcode.NewQRCodeReader()
	result, err := codeReader.Decode(bmp, nil)
	if err != nil {
		return err
	}
	return c.JSON(result.GetText())
}
