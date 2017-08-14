// mirth-api project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/caimeo/iniflags"
	"github.com/caimeo/stickyjar/tracer"
)

//“https://localhost:443/mirth/api/3.5.0/”
var remoteServer = flag.String("server", "localhost", "The remote server name")
var remotePort = flag.String("port", "443", "The remote server port")
var remoteAPIVersion = flag.String("api_version", "3.5.0", "The remote Mirth server API version")
var remoteUsername = flag.String("username", "", "The remote user name")
var remotePassword = flag.String("password", "", "The remote password")
var tlsVerify = flag.Bool("tls", true, "Is TLS verfify on")
var verboseMode = flag.Bool("verbose", false, "Verbose console output.")

var t tracer.Tracer

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	go monitorErrors(f.CommonErrorChannel())
	t := tracer.New(*verboseMode)

	t.Always("Mirth API")
	f.Tracer = *t

	mirthagent.TLSVerify = *tlsVerify
	a := mirthagent.New(*remoteServer, *remotePort)

	//login, _, restoreable := m.LoginStatus()

	var connected bool
	//if !login && !restoreable {

	cch, _ := a.Login(*remoteUsername, *remotePassword)

	connected, ok := <-cch
	//}

	//ok := <-c

	if connected {
		fmt.Println("IS OK", ok)
	}
	fmt.Println(a.LoginStatus())

	chc, _ := a.Connect()
	ok2, ok := <-chc

	fmt.Println(ok2)

	fmt.Println(a.LoginStatus())

	r, _ := a.API.System.Info() // .SystemInfo()

	si := <-r

	q, _ := a.API.Channel.Status()

	cs := <-q

	time.Sleep(1 * time.Second)
	fmt.Println(si)
	fmt.Println(cs)

	for i, v := range cs {
		fmt.Println(i, v)
	}
}

func monitorErrors(e chan error) {
	for err := range e {
		x, ok := interface{}(err).(model.ExtendedError)
		if ok {
			fmt.Fprintln(os.Stderr, "EXTENDED ERROR\n", err, "\n", x.Cause())
		} else {
			fmt.Fprintln(os.Stderr, "ERROR\n", err)
		}
	}
}
