package system

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
)

func (Ω *System) Info() (chan model.SystemInfo, chan error) {
	c := make(chan model.SystemInfo, 1)
	ec := make(chan error, 1)
	go Ω.info(c, ec)
	return c, ec
}

func (Ω *System) info(c chan model.SystemInfo, ec chan error) {
	defer close(c)
	defer close(ec)
	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.System.Info())
	r, b, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "System info could not be retrieved") {
		return
	}
	c <- model.SystemInfoFromXml(b)
}
