package mirthagent

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

func (a *Agent) loginPath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", a.Server, a.Port)
}

func (a *Agent) loginResp(resp gorequest.Response, body string, errs []error) {
	Tracer.Verbose(strconv.Itoa(resp.StatusCode))
	Tracer.Verbose(body)
}

func (a *Agent) Login(username string, password string) {
	a.request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	a.request.Type("form-data")
	a.request.Post(a.loginPath())
	a.request.Send(fmt.Sprintf("username=%s", username))
	a.request.Send(fmt.Sprintf("password=%s", password))

	traceCurl(a.request)

	a.request.End(a.loginResp)
}
