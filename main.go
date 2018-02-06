// mirth-api project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/caimeo/console"
	"github.com/caimeo/iniflags"
)


var remoteServer = flag.String("server", "localhost", "The remote server name")
var remotePort = flag.String("port", "40443", "The remote server port")
var remoteAPIVersion = flag.String("api_version", "3.5.0", "The remote Mirth server API version")
var remoteUsername = flag.String("username", "", "The remote user name")
var remotePassword = flag.String("password", "", "The remote password")
var tlsVerify = flag.Bool("tls", true, "Is TLS verfify on")
var verboseMode = flag.Bool("verbose", false, "Verbose console output.")
var debugMode = flag.Bool("debug", false, "Debug console output.")

var con *console.Console

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	//setup output
	go monitorErrors(errors.CommonErrorChannel())
	con := console.Init(*verboseMode, *debugMode)
	errors.Console = con
	con.Always("Mirth API")

	//setup MirthAgent
	mirthagent.TLSVerify = *tlsVerify
	a := mirthagent.New(*remoteServer, *remotePort, *remoteAPIVersion)

	//get the current status
	_, _, isRestoreable := a.LoginStatus()

	isConnected := false

	if isRestoreable {
		connectCh, _ := a.Connect()
		isConnected, _ = <-connectCh
	}

	if isConnected {
		con.Verbose("Connection Restored: ", isConnected)
	} else {
		if len(*remoteUsername) > 0 && len(*remotePassword) > 0 {
			loginCh, _ := a.Login(*remoteUsername, *remotePassword)
			isConnected, _ = <-loginCh
			if isConnected {
				con.Verbose("Connection New: ", isConnected)
			}
		} else {
			con.Always("Cannot connect, user name and password required.")
			os.Exit(1)
		}
	}

	if isConnected {
		con.Verbose("Ready")
		con.Debug(a.LoginStatus())
	} else {
		con.Always("Cannot connect")
		os.Exit(2)
	}

	rc, _ := a.API.Server.ResourcesList()

	r := <-rc

	//	cg := model.ChannelGroup{Versionø: "3.5.0", Idø: "b49f7831-d524-4a02-bb8c-61db3921166b", Nameø: "TEST Name", Descriptionø: "This is the description"}

	//	c1 := model.ChannelGroupChannel{Versionø: "3.5.0", Idø: "abc"}
	//	c2 := model.ChannelGroupChannel{Versionø: "3.5.0", Idø: "xyz"}

	//	cg.AppendChannel(c1)
	//	cg.AppendChannel(c2)
	//	cg.SetName("A different name")

	//	cg.SetChannels(append(cg.Channels(), c2))

	//	console.Always(cg)

	//x, _ := xml.Marshal(cg)
	console.Always(len(r))

	x := r[0].Id
	console.Always("- - - - -")
	console.Always(string(x))
	console.Always("---------")

	fc, _ := a.API.Server.SetDefaultResourceDirectory("/mnt/mirth/data/java")

	f := <-fc

	console.Always("xxx- - - - -")
	console.Always(f)
	console.Always("---------")

	//	time.Sleep(5 * time.Second)

}

func monitorErrors(e chan error) {
	for err := range e {
		x, ok := interface{}(err).(errors.ExtendedError)
		if ok {
			fmt.Fprintln(os.Stderr, "EXTENDED ERROR\n", err, "\n", x.Cause())
		} else {
			fmt.Fprintln(os.Stderr, "ERROR\n", err)
		}
	}
}
