package main

import (
	"encoding/base64"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"log"
)

// preLoadConfig creates an empty config if none exists
func preLoadConfig() {
	m, err := sto.ReadConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if m != nil {
		return
	}
	_, err = sto.UpsertConfig(&storage.Config{
		Version:        "",
		WsUrl:          "",
		Realm:          "",
		ConsumerKey:    "",
		ConsumerSecret: "",
		TokenID:        "",
		TokenSecret:    "",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// prePlugin creates an empty plugin if none exists
func prePlugin() {
	m, err := sto.ReadPlugins()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if m != nil {
		return
	}
	var x []storage.Plugin
	x = append(x, storage.Plugin{
		Name:        "OCR Reader",
		Description: "Reads text from an image",
		Enabled:     true,
		Mapping: []storage.ProfileMapping{{
			LabelText:   "OCR Result",
			Name:        "ocr-text",
			Value:       "",
			Placeholder: "",
			Hidden:      false,
			Readonly:    true,
		}},
	})
	x = append(x, storage.Plugin{
		Name:        "Barcode Reader",
		Description: "Reads barcodes as text",
		Enabled:     true,
		Mapping: []storage.ProfileMapping{{
			LabelText:   "Barcode Result",
			Name:        "barcode-text",
			Value:       "",
			Placeholder: "",
			Hidden:      false,
			Readonly:    true,
		}},
	})
	_, err = sto.UpsertPlugins(x)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// preProfile creates an empty profile if none exists
func preProfile() {
	m, err := sto.ReadProfiles()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if m != nil {
		return
	}
	var x []storage.Profile
	x = append(x, storage.Profile{
		Name: "Upload a File",
		Template: `<add xmlns="urn:messages_{{.Version}}.platform.webservices.netsuite.com">
					<record xsi:type="ns8:File" xmlns:ns8="urn:filecabinet_{{.Version}}.documents.webservices.netsuite.com">
						<ns8:name xsi:type="xsd:string">{{.FileName}}</ns8:name>
						<ns8:attachFrom xsi:type="ns9:FileAttachFrom" xmlns:ns9="urn:types.filecabinet_{{.Version}}.documents.webservices.netsuite.com">_computer</ns8:attachFrom>
						<ns8:content>{{.FileContent}}</ns8:content>
						<ns8:folder internalId="{{.folder}}" xsi:type="ns10:RecordRef" xmlns:ns10="urn:core_{{.Version}}.platform.webservices.netsuite.com"/>
						<ns8:description xsi:type="xsd:string">{{.Ocr}}</ns8:description>
					</record>
				</add>`,
		Mapping: []storage.ProfileMapping{{
			LabelText:   "NetSuite Folder",
			Name:        "folder",
			Value:       "-10",
			Placeholder: "Set Netsuite Folder",
			Hidden:      false,
			Readonly:    false,
		}},
		Plugin: "OCR Reader",
	})
	_, err = sto.UpsertProfiles(x)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// preLoadAdminUser creates a default admin user if not existent
func preLoadAdminUser() {
	m, err := sto.ReadUsers()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if len(m) != 0 {
		return
	}
	key := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	user := storage.User{
		Name:         "admin",
		Username:     "admin",
		Key:          key,
		Enabled:      true,
		Admin:        true,
		AccessToken:  "",
		RefreshToken: "",
		TokenUuid:    "",
		RefreshUuid:  "",
		AtExpires:    0,
		RtExpires:    0,
	}
	err = tok.ApplyTokenValues(key, &user)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var users []storage.User
	users = append(users, user)
	_, err = sto.UpsertUsers(users)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
