package mirthagent

import (
	"net/http"
	"os"
	"path"

	"github.com/caimeo/stickyjar/curljar"
	"github.com/parnurzeal/gorequest"
)

const defaultCookieFile = "cookies.cook"

type Agent struct {
	Server     string
	Port       string
	CookieFile string
	Jar        http.CookieJar
	request    *gorequest.SuperAgent
}

func New(server string, port string) *Agent {
	a := Agent{Server: server, Port: port}

	a.Request()
	return &a
}

func (a *Agent) Request() *gorequest.SuperAgent {
	if a.request == nil {
		a.request = a.newRequest()
	}
	return a.request
}

func (a *Agent) newRequest() *gorequest.SuperAgent {
	r := gorequest.New()
	if a.Jar == nil {
		jar, err := curljar.New(a.cookieFile(), nil)
		check(err)
		a.Jar = jar
	}
	r.Client.Jar = a.Jar
	return r
}

func (a *Agent) cookieFile() string {
	if a.CookieFile == "" {
		ex, err := os.Executable()
		check(err)
		dir := path.Dir(ex)
		a.CookieFile = path.Join(dir, defaultCookieFile)
		Tracer.Verbose(a.CookieFile)
	}
	return a.CookieFile
}
