package channel

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"

	"github.com/parnurzeal/gorequest"
)

func (Ω *Channel) Status() (chan model.ChannelStatusSlice, chan error) {
	c := make(chan model.ChannelStatusSlice, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.Channels.Statuses())
	go status(req, c, ec)
	return c, ec
}

func status(req *gorequest.SuperAgent, c chan model.ChannelStatusSlice, ec chan error) {
	defer close(c)
	defer close(ec)
	r, b, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Channel status could not be retrieved") {
		return
	}
	slice := model.ChannelStatusSlice{Slice: model.ChannelStatusFromXml(b)}
	c <- slice
}
