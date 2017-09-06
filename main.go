// mirth-api project main.go
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
	"github.com/caimeo/console"
	"github.com/caimeo/iniflags"
)

//“https://localhost:443/mirth/api/3.5.0/”
var remoteServer = flag.String("server", "localhost", "The remote server name")
var remotePort = flag.String("port", "443", "The remote server port")
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
	a := mirthagent.New(*remoteServer, *remotePort)

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
		fmt.Println(a.LoginStatus())
	} else {
		con.Always("Cannot connect")
		os.Exit(2)
	}

	//get system info
	//	siCh, _ := a.API.System.Info()
	//	si := <-siCh

	//get channel status
	//	csCh, _ := a.API.Channel.Status()
	//	csSlice := <-csCh

	//	cs := csSlice.Slice

	//output the system info
	//	fmt.Println(si)
	//	fmt.Println(cs)

	//	for i, v := range cs {
	//		fmt.Println(i, v)
	//	}

	//	for i, v := range cs {
	//		ceCh, _ := a.API.Channel.Disable(v.ChannelId)
	//		ce2 := <-ceCh

	//		fmt.Println(i, ce2)
	//	}

	//	ceCh, _ := a.API.Channel.Enable("")
	//	ce2 := <-ceCh

	//	fmt.Println(ce2)

	//	ce3 := csSlice.ToMap()

	//	fmt.Println(ce3)

	cg := model.ChannelGroup{Versionø: "3.5.0", Idø: "b49f7831-d524-4a02-bb8c-61db3921166b", Nameø: "TEST Name", Descriptionø: "This is the description"}

	c1 := model.ChannelGroupChannel{Versionø: "3.5.0", Idø: "abc"}
	c2 := model.ChannelGroupChannel{Versionø: "3.5.0", Idø: "xyz"}
	//	//cg.SetChannels(append(cg.Channels(), c1))
	//cg.SetChannels(append(cg.Channels(), c2))
	//	channels := cg.Channels()
	cg.AppendChannel(c1)
	cg.AppendChannel(c2)
	cg.SetName("A different name")

	cg.SetChannels(append(cg.Channels(), c2))
	//cg.SetChannels(*(append(*(cg.Channels()), c1)))
	//onsole.Always(ch)
	//cg.SetChannels(&ch)
	//	channels = append(channels, c1)
	//	channels = append(channels, c2)
	//	cg.SetChannels(channels)

	console.Always(cg)

	x, _ := xml.Marshal(cg)

	console.Always(string(x))

	time.Sleep(5 * time.Second)

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
