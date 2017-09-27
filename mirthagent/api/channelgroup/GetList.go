package channelgroup

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"

	"github.com/parnurzeal/gorequest"
)

func (Ω *ChannelGroup) GetList() (chan []model.ChannelGroup, chan error) {
	c := make(chan []model.ChannelGroup, 1)
	ec := make(chan error, 1)

	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.ChannelGroups.GetList())
	go getlist(req, c, ec)
	return c, ec
}

func getlist(req *gorequest.SuperAgent, c chan []model.ChannelGroup, ec chan error) {
	defer close(c)
	defer close(ec)
	r, b, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Channel group list could not be retrieved") {
		return
	}
	groups := model.ChannelGroupsFromXml(b)
	c <- groups
}
