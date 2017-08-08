package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/handle"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/parse"

	"github.com/parnurzeal/gorequest"
)

func (a *Agent) SystemInfo(onErr handle.Error, onData handle.SystemInfo) {
	a.request.Get(a.infoPath())
	f := func(r gorequest.Response, b []byte, e []error) {
		a.systemInfoResponder(onErr, onData, r, b, e)
	}
	a.request.EndBytes(f)
}

func (a *Agent) systemInfoResponder(onErr handle.Error, onData handle.SystemInfo, resp gorequest.Response, body []byte, e []error) {
	if errorCheck(onErr, e, "System info could not be retrieved") {
		return
	}
	if resp.StatusCode == 200 {
		onData(parse.SystemInfo(body))
	}
}
