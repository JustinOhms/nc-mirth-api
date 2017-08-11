package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/systeminfo"
)

func (a *Agent) SystemInfo() (chan systeminfo.SystemInfo, chan error) {
	c := make(chan systeminfo.SystemInfo, 1)
	ec := make(chan error, 1)

	go a.systemInfo(c, ec)

	return c, ec
}

func (a *Agent) systemInfo(c chan systeminfo.SystemInfo, ec chan error) {
	defer close(c)
	defer close(ec)

	a.request.Get(a.Paths.System.Info())
	r, b, e := a.request.EndBytes()

	if responseOrStatusErrors(ec, r, e, "System info could not be retrieved") {
		return
	}

	c <- systeminfo.XmlParse(b)

}
