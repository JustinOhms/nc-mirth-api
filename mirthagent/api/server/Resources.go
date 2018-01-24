package server

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/caimeo/console"
	"github.com/parnurzeal/gorequest"
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

func (Ω *Server) ResourcesList() (chan []model.Resource, chan error) {
	c := make(chan []model.Resource, 1)
	ec := make(chan error, 1)

	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.Server.Resources())

	go Ω.resourceList(req, c, ec)

	return c, ec
}

func (Ω *Server) resourceList(req *gorequest.SuperAgent, c chan []model.Resource, ec chan error) {
	defer close(c)
	defer close(ec)
	r, b, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error reading resource list") {
		return
	}
	list := model.ResourcesFromXml(b)
	console.Always(list)
	c <- list
}

// warning only use this if there is only the default
func (Ω *Server) SetDefaultResourceDirectory(dir string) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)

	defRes := model.NewDefaultResource()
	defRes.Directory = dir
	r := make([]model.Resource, 0)
	r = append(r, model.Resource(*defRes))
	xml, _ := model.ResourcesToXml(r)

	req := Ω.Session.NewRequest().Put(Ω.Session.Paths.Server.Resources())
	req.Type("xml")
	req.RawString = string(xml)
	req.BounceToRawString = true

	go Ω.setDefaultResourceDirectory(req, c, ec)

	return c, ec
}

func (Ω *Server) setDefaultResourceDirectory(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error setting default resource directory") {
		return
	}
	c <- true
}
