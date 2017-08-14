package channel

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
)

func (Ω *Channel) Status() (chan []model.ChannelStatus, chan error) {
	c := make(chan []model.ChannelStatus, 1)
	ec := make(chan error, 1)
	go Ω.status(c, ec)
	return c, ec
}

func (Ω *Channel) status(c chan []model.ChannelStatus, ec chan error) {
	defer close(c)
	defer close(ec)
	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.Channels.Statuses())
	r, b, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "Channel status could not be retrieved") {
		return
	}
	c <- model.ChannelStatusFromXml(b)
}
