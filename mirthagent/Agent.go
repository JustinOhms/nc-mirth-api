package mirthagent

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/NavigatingCancer/mirth-api/mirthagent/resource"
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
	Paths       resource.Paths
}

var TLSVerify bool = true

func New(server string, port string) *Agent {
	a := Agent{Server: server, Port: port}
	a.Paths = resource.New(server, port)
	a.TLSVerify = TLSVerify
	a.loginStatus = false
	a.Request()

	return &a
}

//returns the default request object
func (a *Agent) Request() *gorequest.SuperAgent {
	if a.request == nil {
		a.request = a.NewRequest()
	}

	return a.request
}

//returns a new request object clone session from default request
func (a *Agent) NewRequest() *gorequest.SuperAgent {
	r := gorequest.New()

	r.TLSClientConfig(&tls.Config{InsecureSkipVerify: !a.TLSVerify})

	//if no jar is initialized set the jar
	if a.Jar == nil {
		jar, err := curljar.New(a.cookieFile(), nil)
		checkErrorAndPanic(err)
		a.Jar = jar
	}
	r.Client.Jar = a.Jar

	//if we are not logged in and we have a restorable cookie file restore it
	if a.loginStatus == false && a.restorableSession() {
		fmt.Println("RESTOREING")
		(a.Jar).(restorable.Restorable).Restore()
	}

	return r
}
