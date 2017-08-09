// mirth-api project main.go
package main

import (
	"flag"
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/channelstatus"
	"github.com/NavigatingCancer/mirth-api/mirthagent/model/systeminfo"
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

	t := tracer.New(*verboseMode)

	t.Always("Mirth API")
	mirthagent.Tracer = *t

	mirthagent.TLSVerify = *tlsVerify
	m := mirthagent.New(*remoteServer, *remotePort)

	login, _, restoreable := m.LoginStatus()

	if !login && !restoreable {
		m.Login(*remoteUsername, *remotePassword)
	}

	fmt.Println(m.LoginStatus())

	m.Connect()

	fmt.Println(m.LoginStatus())

	//m.SystemInfo(handleError, infoResponse)

	//m.ChannelStatus(handleError, statResponse)

}

func infoResponse(i systeminfo.SystemInfo) {
	fmt.Println("INFORESPONSE")
	fmt.Println(i)
}

func handleError(e error) {
	fmt.Println("ERRROR")
	fmt.Println(e.Error())
}

func statResponse(i []channelstatus.ChannelStatus) {
	fmt.Println("CHANNELSTATUS")
	fmt.Println(i)
}
