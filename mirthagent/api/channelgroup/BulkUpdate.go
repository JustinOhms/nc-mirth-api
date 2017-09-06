package channelgroup

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"

	"github.com/parnurzeal/gorequest"
)

func (Ω *ChannelGroup) BulkUpdate() (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.ChannelGroups.BulkUpdate())
	go bulkupdate(req, c, ec)
	return c, ec
}

func bulkupdate(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, b, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Channel Groups could not be set") {
		return
	}
	//slice := model.ChannelStatusSlice{Slice: model.ChannelStatusFromXml(b)}
	//c <- slice
}
