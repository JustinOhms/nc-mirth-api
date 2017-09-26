package server

import (
	"github.com/parnurzeal/gorequest"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
)

func (Ω *Server) ResourceReload(resourceId string) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)

	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.Server.ResourceReload(resourceId))

	go Ω.resourceReload(req, c, ec)

	return c, ec
}

func (Ω *Server) resourceReload(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error reloading resource") {
		return
	}
	c <- true
}
