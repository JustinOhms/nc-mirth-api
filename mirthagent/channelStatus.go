package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/channelstatus"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/errorhandler"

	"github.com/parnurzeal/gorequest"
)

func (a *Agent) ChannelStatus(onErr errorhandler.Handler, onData channelstatus.Handler) {
	a.request.Get(a.channelStatusPath())
	f := func(r gorequest.Response, b []byte, e []error) {
		a.channelStatusHandler(onErr, onData, r, b, e)
	}
	a.request.EndBytes(f)
}

func (a *Agent) channelStatusHandler(onErr errorhandler.Handler, onData channelstatus.Handler, resp gorequest.Response, body []byte, e []error) {
	if errorCheck(onErr, e, "Channel status could not be retrieved") {
		return
	}
	if resp.StatusCode == 200 {
		onData(channelstatus.XmlParse(body))
	}
}
