package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/channelstatus"
)

func (a *Agent) ChannelStatus() (chan []channelstatus.ChannelStatus, chan error) {
	c := make(chan []channelstatus.ChannelStatus, 1)
	ec := make(chan error, 1)

	go a.channelstatus(c, ec)

	return c, ec
}

func (a *Agent) channelstatus(c chan []channelstatus.ChannelStatus, ec chan error) {
	defer close(c)
	defer close(ec)

	a.request.Get(a.Paths.Channels.Statuses())
	r, b, e := a.request.EndBytes()

	if responseOrStatusErrors(ec, r, e, "Channel status could not be retrieved") {
		return
	}

	c <- channelstatus.XmlParse(b)

}
