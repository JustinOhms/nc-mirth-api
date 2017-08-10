package mirthagent

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/NavigatingCancer/mirth-api/mirthagent/model/user"
	"github.com/caimeo/stickyjar/restorable"
)

func (a *Agent) Login(username string, password string) chan bool {
	c := make(chan bool, 1)

	a.userName = username

	go a.login(c, username, password)

	return c
}

func (a *Agent) login(c chan bool, username string, password string) {
	a.userName = username
	a.password = password

	a.request.Type("form-data")
	a.request.Post(a.Paths.Users.Login())
	a.request.Send(fmt.Sprintf("username=%s", a.userName))
	a.request.Send(fmt.Sprintf("password=%s", a.password))

	traceCurl(a.request)

	r, b, _ := a.request.End()

	Tracer.Verbose(strconv.Itoa(r.StatusCode))
	Tracer.Verbose(b)
	if r.StatusCode == 200 {
		a.loginStatus = true
		c <- true
	} else {
		a.loginStatus = false
		c <- false
	}
	close(c)
}

func (a *Agent) Connect() chan bool {
	c := make(chan bool, 1)
	if a.loginStatus == false && a.restorableSession() {
		go a.connect(c)
	} else {
		c <- false
	}
	return c
}

func (a *Agent) connect(c chan bool) {
	a.request.Get(a.Paths.Users.Login())
	r, b, _ := a.request.EndBytes()

	Tracer.Verbose(strconv.Itoa(r.StatusCode))
	Tracer.Verbose(string(b[:]))
	if r.StatusCode == 200 {
		u := user.XmlParse(b)
		a.userName = u.UserName
		a.loginStatus = true
		c <- true
	} else {
		//if connect fails we clear the cookie file
		os.Remove(a.cookieFile())
		a.loginStatus = false
		c <- false
	}
	close(c)
}

func (a *Agent) LoginStatus() (loggedIn bool, userName string, restorable bool) {
	return a.loginStatus, a.userName, a.restorableSession()
}

func (a *Agent) cookieFile() string {
	if a.CookieFile == "" {
		ex, err := os.Executable()
		checkErrorAndPanic(err)
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
	return ok && a.hasCookieFile()
}
