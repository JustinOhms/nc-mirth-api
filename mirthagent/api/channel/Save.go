package channel

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Channel) Save(channelId string, channelXmlText []byte) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	url := Ω.Session.Paths.Channels.Save(channelId)

	req := Ω.Session.NewRequest().Put(url)
	req.Type("xml")
	req.RawString = string(channelXmlText)
	req.BounceToRawString = true

	go save(req, c, ec)
	return c, ec
}

func save(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if errors.ResponseOrStatusErrors(ec, r, e, "Error saving channel(s)") {
		return
	}
	c <- true
}
