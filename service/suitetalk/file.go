package suitetalk

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"text/template"
)

type Response struct {
	Body struct {
		Text        string `xml:",chardata"`
		AddResponse struct {
			WriteResponse struct {
				BaseRef struct {
					InternalId string `xml:"internalId,attr"`
				} `xml:"baseRef"`
			} `xml:"writeResponse"`
		} `xml:"addResponse"`
		Fault struct {
			Faultstring string `xml:"faultstring"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type File interface {
	BuildTemplate(settings Settings) ([]byte, error)
	ParseResponse(body []byte) (*Response, error)
}

type file struct {
	Template    string
	FileName    string
	FileContent string
}

func NewFile(template string, filename string, content *bytes.Buffer) File {
	f := &file{
		template,
		filename,
		base64.StdEncoding.EncodeToString(content.Bytes()),
	}
	return f
}

// BuildTemplate build xml template for upload
func (f *file) BuildTemplate(settings Settings) ([]byte, error) {
	type upload struct {
		Version     string
		Realm       string
		ConsumerKey string
		TokenID     string
		Nonce       string
		Timestamp   string
		Signature   string
		FileName    string
		FileContent string
	}
	nonce := settings.GetNonce()
	timestamp := settings.GetTimestamp()
	u := upload{
		Version:     settings.GetVersion(),
		Realm:       settings.GetRealm(),
		ConsumerKey: settings.GetConsumerKey(),
		TokenID:     settings.GetTokenID(),
		Nonce:       nonce,
		Timestamp:   timestamp,
		Signature:   settings.GetSignature(nonce, timestamp),
		FileName:    f.FileName,
		FileContent: f.FileContent,
	}

	envelope := `<?xml version="1.0" encoding="UTF-8"?>
			<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
			<soapenv:Header>
				<ns1:tokenPassport soapenv:actor="http://schemas.xmlsoap.org/soap/actor/next" soapenv:mustUnderstand="0" xmlns:ns1="urn:messages_{{.Version}}.platform.webservices.netsuite.com">
					<ns2:account xmlns:ns2="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.Realm}}</ns2:account>
					<ns3:consumerKey xmlns:ns3="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.ConsumerKey}}</ns3:consumerKey>
					<ns4:token xmlns:ns4="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.TokenID}}</ns4:token>
					<ns5:nonce xmlns:ns5="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.Nonce}}</ns5:nonce>
					<ns6:timestamp xmlns:ns6="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.Timestamp}}</ns6:timestamp>
					<ns7:signature algorithm="HMAC-SHA256" xmlns:ns7="urn:core_{{.Version}}.platform.webservices.netsuite.com">{{.Signature}}</ns7:signature>
				</ns1:tokenPassport>
			</soapenv:Header>
			<soapenv:Body>`
	envelope += f.Template
	envelope += `</soapenv:Body>
		</soapenv:Envelope>`

	tmpl, err := template.New("upload").Parse(envelope)
	if err != nil {
		return nil, err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, u)
	if err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}

func (f *file) ParseResponse(body []byte) (*Response, error) {
	e := &Response{}
	err := xml.Unmarshal(body, e)
	if err != nil {
		return e, err
	}
	return e, nil
}
