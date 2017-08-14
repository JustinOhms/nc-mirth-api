package mirthagent

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/api"
	"github.com/NavigatingCancer/mirth-api/mirthagent/session"
	"github.com/parnurzeal/gorequest"
)

const defaultCookieFile = "cookies.cook"

type Agent struct {
	session *session.Session
	API     *api.API
}

var TLSVerify bool = true

func New(server string, port string) *Agent {
	session.TLSVerify = TLSVerify
	a := Agent{}
	a.session = session.New(server, port)
	a.API = api.New(a.session)
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

func (Ω *Agent) Connect() (chan bool, chan error) {
	return Ω.session.Connect()
}

func (Ω *Agent) LoginStatus() (loggedIn bool, userName string, restorable bool) {
	return Ω.session.LoginStatus()
}
