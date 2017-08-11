package api

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/channelstatus"
)

type API struct {
	Channels channels
	System   system
	agent    mirthagent.Agent
}

func New(a *requestagent.Agent) API {
	api := API{agent: a}
	api.Channels = channels{a: *a}
	api.System = system{a: *a}
	return api
}

type channels struct {
	a requestagent.Agent
}

type system struct {
	a requestagent.Agent
}

func (Ω channels) ChannelStatus() (chan []channelstatus.ChannelStatus, chan error) {
	c := make(chan []channelstatus.ChannelStatus, 1)
	ec := make(chan error, 1)

	go Ω.channelstatus(c, ec)

	return c, ec
}

func (Ω *channels) channelstatus(c chan []channelstatus.ChannelStatus, ec chan error) {
	defer close(c)
	defer close(ec)
	req := Ω.a.NewRequest().Get(Ω.a.Paths().Channels.Statuses())
	r, b, e := req.EndBytes()
	if responseOrStatusErrors(ec, r, e, "Channel status could not be retrieved") {
		return
	}
	c <- channelstatus.XmlParse(b)
}
