package channel

import (
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Channel) Enable(channelId string) (chan bool, chan error) {
	return Ω.setEnable(channelId, true)

}

func (Ω *Channel) Disable(channelId string) (chan bool, chan error) {
	return Ω.setEnable(channelId, false)
}

func (Ω *Channel) setEnable(channelId string, on bool) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.Channels.SetEnable())
	req.Send(fmt.Sprintf("channelId=%s", channelId))
	req.Send(fmt.Sprintf("enabled=%t", on))
	go setEnable(req, c, ec)
	return c, ec
}

func setEnable(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "Error enableing channel") {
		return
	}
	c <- true
}
