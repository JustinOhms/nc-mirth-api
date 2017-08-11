package mirthagent

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/NavigatingCancer/mirth-api/mirthagent/model/user"
	"github.com/caimeo/stickyjar/restorable"
)

func (Ω *Agent) Login(username string, password string) chan bool {
	c := make(chan bool, 1)

	Ω.userName = username

	go Ω.login(c, username, password)

	return c
}

func (Ω *Agent) login(c chan bool, username string, password string) {
	Ω.userName = username
	Ω.password = password

	Ω.request.Type("form-data")
	Ω.request.Post(Ω.Paths.Users.Login())
	Ω.request.Send(fmt.Sprintf("username=%s", Ω.userName))
	Ω.request.Send(fmt.Sprintf("password=%s", Ω.password))

	traceCurl(Ω.request)

	r, b, _ := Ω.request.End()

	Tracer.Verbose(strconv.Itoa(r.StatusCode))
	Tracer.Verbose(b)
	if r.StatusCode == 200 {
		Ω.loginStatus = true
		c <- true
	} else {
		Ω.loginStatus = false
		c <- false
	}
	close(c)
}

func (Ω *Agent) Connect() chan bool {
	c := make(chan bool, 1)
	if Ω.loginStatus == false && Ω.restorableSession() {
		go Ω.connect(c)
	} else {
		c <- false
	}
	return c
}

func (Ω *Agent) connect(c chan bool) {
	Ω.request.Get(Ω.Paths.Users.Login())
	r, b, _ := Ω.request.EndBytes()

	Tracer.Verbose(strconv.Itoa(r.StatusCode))
	Tracer.Verbose(string(b[:]))
	if r.StatusCode == 200 {
		u := user.XmlParse(b)
		Ω.userName = u.UserName
		Ω.loginStatus = true
		c <- true
	} else {
		//if connect fails we clear the cookie file
		os.Remove(Ω.cookieFile())
		Ω.loginStatus = false
		c <- false
	}
	close(c)
}

func (Ω *Agent) LoginStatus() (loggedIn bool, userName string, restorable bool) {
	return Ω.loginStatus, Ω.userName, Ω.restorableSession()
}

func (Ω *Agent) cookieFile() string {
	if Ω.CookieFile == "" {
		ex, err := os.Executable()
		checkErrorAndPanic(err)
		dir := path.Dir(ex)
		Ω.CookieFile = path.Join(dir, defaultCookieFile)
		Tracer.Verbose(Ω.CookieFile)
	}
	return Ω.CookieFile
}

func (Ω *Agent) hasCookieFile() bool {
	_, err := os.Stat(Ω.cookieFile())
	return (err == nil)
}

func (Ω *Agent) restorableSession() bool {
	_, ok := interface{}(Ω.Jar).(restorable.Restorable)
	return ok && Ω.hasCookieFile()
}
