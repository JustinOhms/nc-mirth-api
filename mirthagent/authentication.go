package mirthagent

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

func (a *Agent) loginPath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", a.Server, a.Port)
}

func (a *Agent) loginResp(resp gorequest.Response, body string, errs []error) {
	Tracer.Verbose(strconv.Itoa(resp.StatusCode))
	Tracer.Verbose(body)
	if resp.StatusCode == 200 {
		a.loginStatus = true
	}
}

func (a *Agent) Login(username string, password string) {
	a.userName = username
	a.password = password

	a.request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	a.request.Type("form-data")
	a.request.Post(a.loginPath())
	a.request.Send(fmt.Sprintf("username=%s", a.userName))
	a.request.Send(fmt.Sprintf("password=%s", a.password))

	traceCurl(a.request)

	a.request.End(a.loginResp)
}

func (a *Agent) LoginStatus() (loggedIn bool, userName string, cookie bool) {
	_, err := os.Stat(a.cookieFile())
	return a.loginStatus, a.userName, (err == nil)
}

func (a *Agent) Reconnect() {

}
