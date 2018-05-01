package channel

import (
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Channel) RedeployAll() (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.Channels.RedeployAll())

	req.Query(fmt.Sprintf("returnErrors=false"))

	go redeployAll(req, c, ec)
	return c, ec
}

func redeployAll(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error redeploying all channels") {
		return
	}
	c <- true
}
