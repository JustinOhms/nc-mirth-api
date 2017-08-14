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

	//setup output
	go monitorErrors(f.CommonErrorChannel())
	t := tracer.New(*verboseMode)
	f.Tracer = *t
	t.Always("Mirth API")

	//setup MirthAgent
	mirthagent.TLSVerify = *tlsVerify
	a := mirthagent.New(*remoteServer, *remotePort)

	//get the current status
	_, _, isRestoreable := a.LoginStatus()

	isConnected := false

	if isRestoreable {
		connectCh, _ := a.Connect()
		isConnected, _ = <-connectCh
	}

	if isConnected {
		t.Verbose("Connection Restored: ", isConnected)
	} else {
		if len(*remoteUsername) > 0 && len(*remotePassword) > 0 {
			loginCh, _ := a.Login(*remoteUsername, *remotePassword)
			isConnected, _ = <-loginCh
			if isConnected {
				t.Verbose("Connection New: ", isConnected)
			}
		} else {
			t.Always("Cannot connect, user name and password required.")
			os.Exit(1)
		}
	}

	if isConnected {
		t.Verbose("Ready")
		fmt.Println(a.LoginStatus())
	} else {
		t.Always("Cannot connect")
		os.Exit(2)
	}

	//get system info
	r, _ := a.API.System.Info()
	si := <-r

	//get channel status
	q, _ := a.API.Channel.Status()
	cs := <-q

	//wait a bit for channels to clear
	time.Sleep(1 * time.Second)

	//output the system info
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
