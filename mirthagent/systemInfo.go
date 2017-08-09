package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/errorhandler"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/systeminfo"

	"github.com/parnurzeal/gorequest"
)

func (a *Agent) SystemInfo(onErr errorhandler.Handler, onData systeminfo.Handler) {
	a.request.Get(a.infoPath())
	f := func(r gorequest.Response, b []byte, e []error) {
		a.systemInfoHandler(onErr, onData, r, b, e)
	}
	a.request.EndBytes(f)
}

func (a *Agent) systemInfoHandler(onErr errorhandler.Handler, onData systeminfo.Handler, resp gorequest.Response, body []byte, e []error) {
	if errorCheck(onErr, e, "System info could not be retrieved") {
		return
	}
	if resp.StatusCode == 200 {
		onData(systeminfo.XmlParse(body))
	}
}
