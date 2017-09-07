package channelgroup

import (
	"encoding/xml"
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/parnurzeal/gorequest"
)

func (Ω *ChannelGroup) SetGroups(cgrps []model.ChannelGroup) (chan bool, chan error) {
	channelGroups, _ := xml.Marshal(cgrps)
	//channelGroups = []byte(strings.Replace(string(channelGroups), "<channel ", "\n<channel ", -1))
	return Ω.BulkUpdate(channelGroups, []byte(""))
}

func (Ω *ChannelGroup) RemoveGroups(groupIds []string) (chan bool, chan error) {
	grouplistxml := ""
	for _, v := range groupIds {
		grouplistxml = fmt.Sprintf("%s<string>%s</string>", grouplistxml, v)
	}

	return Ω.BulkUpdate([]byte(""), []byte(grouplistxml))
}

func (Ω *ChannelGroup) BulkUpdate(channelGroups []byte, removedChannelGroupIds []byte) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.ChannelGroups.BulkUpdate())

	req.Type("multipart")
	req.Set("Accept", "application/xml")
	req.Set("Connection", "Keep-Alive")

	channelGroups = []byte(fmt.Sprintf("<set>%s</set>", channelGroups))
	req.SendFile(channelGroups, "", "channelGroups", "application/xml")

	removedChannelGroupIds = []byte(fmt.Sprintf("<set>%s</set>", removedChannelGroupIds))
	req.SendFile(removedChannelGroupIds, "", "removedChannelGroupIds", "application/xml")

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

	m := simpleBoolean{}
	xml.Unmarshal(b, &m)

	c <- m.Boolean
}

type simpleBoolean struct {
	XMLName xml.Name `xml:"boolean"`
	Boolean bool     `xml:",chardata"`
}
