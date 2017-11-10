package server

import (
	"encoding/xml"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Server) ConfigurationMapSave(config model.ConfigurationMap) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	url := Ω.Session.Paths.Server.ConfigurationMap()

	req := Ω.Session.NewRequest().Put(url)
	req.Type("xml")
	res, _ := xml.Marshal(config)
	req.RawString = string(res)
	req.BounceToRawString = true

	go configurationMapSave(req, c, ec)
	return c, ec
}

func configurationMapSave(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error uploading configuration map") {
		return
	}
	c <- true
}
