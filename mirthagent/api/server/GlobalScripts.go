package server

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Server) GlobalScripts(channelId string, channelXmlText []byte) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	url := Ω.Session.Paths.Server

	req := Ω.Session.NewRequest().Put(url)
	req.Type("xml")
	req.RawString = string(channelXmlText)
	req.BounceToRawString = true

	go save(req, c, ec)
	return c, ec
}

func globalScripts(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error saving channel(s)") {
		return
	}
	c <- true
}
