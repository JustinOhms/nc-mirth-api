package session

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/user"
	"github.com/NavigatingCancer/mirth-api/mirthagent/resource"
	"github.com/caimeo/stickyjar/curljar"
	"github.com/caimeo/stickyjar/restorable"
	"github.com/parnurzeal/gorequest"
)

const defaultCookieFile = "cookies.cook"

type Session struct {
	Server     string
	Port       string
	CookieFile string
	TLSVerify  bool
	Jar        http.CookieJar

	request  *gorequest.SuperAgent
	userName string
	password string

	loginStatus bool

	Paths resource.Paths
}

var TLSVerify bool = true

func New(server string, port string) *Session {
	c := Session{Server: server, Port: port}
	c.Paths = resource.PathsNew(server, port)

	c.TLSVerify = TLSVerify
	c.loginStatus = false
	c.Request()

	return &c
}

//returns the default request object
func (Ω *Session) Request() *gorequest.SuperAgent {
	if Ω.request == nil {
		Ω.request = Ω.NewRequest()
	}

	return Ω.request
}

//returns a new request object clone session from default request
func (Ω *Session) NewRequest() *gorequest.SuperAgent {
	r := gorequest.New()

	r.TLSClientConfig(&tls.Config{InsecureSkipVerify: !Ω.TLSVerify})

	//if no jar is initialized set the jar
	if Ω.Jar == nil {
		jar, err := curljar.New(Ω.cookieFile(), nil)
		f.CheckErrorAndPanic(err)
		Ω.Jar = jar
	}
	r.Client.Jar = Ω.Jar

	//if we are not logged in and we have a restorable cookie file restore it
	if Ω.loginStatus == false && Ω.restorableSession() {
		fmt.Println("RESTOREING")
		(Ω.Jar).(restorable.Restorable).Restore()
	}

	return r
}

func (Ω *Session) Login(username string, password string) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)

	Ω.userName = username

	go Ω.login(c, ec, username, password)

	return c, ec
}

func (Ω *Session) login(c chan bool, ec chan error, username string, password string) {
	defer close(c)
	defer close(ec)

	Ω.userName = username
	Ω.password = password

	Ω.request.Type("form-data")
	Ω.request.Post(Ω.Paths.Users.Login())
	Ω.request.Send(fmt.Sprintf("username=%s", Ω.userName))
	Ω.request.Send(fmt.Sprintf("password=%s", Ω.password))

	f.TraceCurl(Ω.request)

	r, b, e := Ω.request.End()

	if f.ResponseOrStatusErrors(ec, r, e, "Login Problem") {
		return
	}

	f.Tracer.Verbose(strconv.Itoa(r.StatusCode))
	f.Tracer.Verbose(b)

	if r.StatusCode == 200 {
		Ω.loginStatus = true
		c <- true
	} else {
		Ω.loginStatus = false
		c <- false
	}

}

func (Ω *Session) Connect() (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)

	if Ω.loginStatus == false && Ω.restorableSession() {
		go Ω.connect(c, ec)
	} else {
		c <- false
	}
	return c, ec
}

func (Ω *Session) connect(c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)

	Ω.request.Get(Ω.Paths.Users.Login())
	r, b, e := Ω.request.EndBytes()

	if f.ResponseOrStatusErrors(ec, r, e, "Connnection Problem") {
		return
	}

	f.Tracer.Verbose(strconv.Itoa(r.StatusCode))
	f.Tracer.Verbose(string(b[:]))

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
}

func (Ω *Session) LoginStatus() (loggedIn bool, userName string, restorable bool) {
	return Ω.loginStatus, Ω.userName, Ω.restorableSession()
}

func (Ω *Session) cookieFile() string {
	if Ω.CookieFile == "" {
		ex, err := os.Executable()
		f.CheckErrorAndPanic(err)
		dir := path.Dir(ex)
		Ω.CookieFile = path.Join(dir, defaultCookieFile)
		f.Tracer.Verbose(Ω.CookieFile)
	}
	return Ω.CookieFile
}

func (Ω *Session) hasCookieFile() bool {
	_, err := os.Stat(Ω.cookieFile())
	return (err == nil)
}

func (Ω *Session) restorableSession() bool {
	_, ok := interface{}(Ω.Jar).(restorable.Restorable)
	return ok && Ω.hasCookieFile()
}