package mirthagent

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/caimeo/stickyjar/curljar"
	"github.com/caimeo/stickyjar/restorable"

	"github.com/parnurzeal/gorequest"
)

const defaultCookieFile = "cookies.cook"

type Agent struct {
	Server      string
	Port        string
	CookieFile  string
	TLSVerify   bool
	Jar         http.CookieJar
	request     *gorequest.SuperAgent
	userName    string
	password    string
	loginStatus bool
}

var TLSVerify bool = true

func New(server string, port string) *Agent {
	a := Agent{Server: server, Port: port}
	a.TLSVerify = TLSVerify
	a.loginStatus = false
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

	r.TLSClientConfig(&tls.Config{InsecureSkipVerify: !a.TLSVerify})

	if a.Jar == nil {
		jar, err := curljar.New(a.cookieFile(), nil)
		check(err)
		a.Jar = jar
	}
	r.Client.Jar = a.Jar

	if a.hasCookieFile() && a.restorableSession() {
		fmt.Println("RESTOREING")
		(a.Jar).(restorable.Restorable).Restore()
	}

	return r
}
