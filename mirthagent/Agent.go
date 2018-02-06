package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/api"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/NavigatingCancer/mirth-api/mirthagent/session"
	"github.com/parnurzeal/gorequest"
)

const defaultCookieFile = "cookies.cook"

type Agent struct {
	session *session.Session
	API     *api.API
}

var TLSVerify bool = true

func New(server string, port string, version string) *Agent {
	session.TLSVerify = TLSVerify
	a := Agent{}
	s := session.New(server, port, version)
	a.session = s
	a.API = api.New(s)
	model.SetAPIVersion(version)
	return &a
}

//returns the default request object
func (Ω *Agent) Request() *gorequest.SuperAgent {
	return Ω.session.Request()
}

//returns a new request object clone session from default request
func (Ω *Agent) NewRequest() *gorequest.SuperAgent {
	return Ω.session.NewRequest()
}

func (Ω *Agent) Login(username string, password string) (chan bool, chan error) {
	return Ω.session.Login(username, password)
}

func (Ω *Agent) Connect() (b chan bool, e chan error) {
	b, e = Ω.session.Connect()
	return
}

func (Ω *Agent) LoginStatus() (loggedIn bool, userName string, restorable bool) {
	return Ω.session.LoginStatus()
}
