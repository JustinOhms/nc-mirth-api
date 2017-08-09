package mirthagent

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/caimeo/stickyjar/restorable"
	"github.com/parnurzeal/gorequest"
)

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

	a.request.Type("form-data")
	a.request.Post(a.loginPath())
	a.request.Send(fmt.Sprintf("username=%s", a.userName))
	a.request.Send(fmt.Sprintf("password=%s", a.password))

	traceCurl(a.request)

	a.request.End(a.loginResp)
}

func (a *Agent) LoginStatus() (loggedIn bool, userName string, cookie bool) {
	return a.loginStatus, a.userName, a.hasCookieFile()
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

func (a *Agent) hasCookieFile() bool {
	_, err := os.Stat(a.cookieFile())
	return (err == nil)
}

func (a *Agent) restorableSession() bool {
	_, ok := interface{}(a.Jar).(restorable.Restorable)
	return ok
}
