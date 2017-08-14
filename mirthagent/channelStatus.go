package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/channelstatus"
)

func (Ω *Agent) ChannelStatus() (chan []channelstatus.ChannelStatus, chan error) {
	c := make(chan []channelstatus.ChannelStatus, 1)
	ec := make(chan error, 1)
	go Ω.channelstatus(c, ec)
	return c, ec
}

func (Ω *Agent) channelstatus(c chan []channelstatus.ChannelStatus, ec chan error) {
	defer close(c)
	defer close(ec)
	req := Ω.NewRequest().Get(Ω.session.Paths.Channels.Statuses())
	r, b, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "Channel status could not be retrieved") {
		return
	}
	c <- channelstatus.XmlParse(b)
}
