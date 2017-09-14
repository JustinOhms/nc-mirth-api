package codeTemplateLibraries

import (
	"encoding/xml"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/parnurzeal/gorequest"
)

func (Ω *CodeTemplateLibraries) BulkUpdate(libraries model.CodeLibraries, templates model.CodeTemplates) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)

	librariesXml, _ := xml.Marshal(libraries)
	templatesXml, _ := xml.Marshal(templates)

	url := Ω.Session.Paths.CodeTemplateLibraries.BulkUpdate()
	req := Ω.Session.NewRequest().Post(url)

	req.Type("multipart")
	req.Set("Accept", "application/xml")
	req.Set("Connection", "Keep-Alive")

	req.SendFile(librariesXml, "", "libraries", "application/xml")
	req.SendFile(templatesXml, "", "updatedCodeTemplates", "application/xml")

	go bulkupdate(req, c, ec)
	return c, ec

}

func bulkupdate(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error saving code template libraries") {
		return
	}
	c <- true
}

type simpleBoolean struct {
	XMLName xml.Name `xml:"boolean"`
	Boolean bool     `xml:",chardata"`
}
