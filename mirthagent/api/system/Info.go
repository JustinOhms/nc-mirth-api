package system

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/NavigatingCancer/mirth-api/mirthagent/ƒ"
	"github.com/parnurzeal/gorequest"
)

func (Ω *System) Info() (chan model.SystemInfo, chan error) {
	c := make(chan model.SystemInfo, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.System.Info())
	go info(req, c, ec)
	return c, ec
}

func info(req *gorequest.SuperAgent, c chan model.SystemInfo, ec chan error) {
	defer close(c)
	defer close(ec)
	r, b, e := req.EndBytes()
	if ƒ.ResponseOrStatusErrors(ec, r, e, "System info could not be retrieved") {
		return
	}
	c <- model.SystemInfoFromXml(b)
}
