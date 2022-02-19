package tesseract

import "github.com/otiai10/gosseract"

type Tesseract interface {
	ReadImage(data []byte) (string, error)
}

type tesseract struct {
}

func NewTesseract() Tesseract {
	return &tesseract{}
}

func (t *tesseract) ReadImage(data []byte) (string, error) {
	client := gosseract.NewClient()
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	err := client.SetImageFromBytes(data)
	if err != nil {
		return "", err
	}
	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
