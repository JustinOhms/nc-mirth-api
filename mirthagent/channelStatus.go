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
	if a.errorCheck(onErr, e, "Channel status could not be retrieved") {
		return
	}

	if a.responseCheck(onErr, resp) {
		onData(channelstatus.XmlParse(body))
	}
}

func (a *Agent) responseCheck(onErr errorhandler.Handler, resp gorequest.Response) bool {
	if resp.StatusCode == 200 {
		return true
	}
	if resp.StatusCode == 401 {
		panic("Unauthorized")
	}

	return false
}
