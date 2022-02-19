package suitetalk

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Upload(settings Settings, file File) (string, error) {

	buildTemplate, err := file.BuildTemplate(settings)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	url := fmt.Sprintf("%s/services/NetSuitePort_%s", settings.GetWsUrl(), settings.GetVersion())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buildTemplate))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "add")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response, err := file.ParseResponse(body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("%s: %s", resp.Status, response.Body.Fault.Faultstring))
	}

	return response.Body.AddResponse.WriteResponse.BaseRef.InternalId, nil
}
