// mirth-api project main.go
package main

import (
	"flag"
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent"
	"github.com/caimeo/iniflags"
	"github.com/caimeo/stickyjar/tracer"
)

//“https://localhost:443/mirth/api/3.5.0/”
var remote_server = flag.String("server", "localhost", "The remote server name")
var remote_port = flag.String("port", "443", "The remote server port")
var remote_API_version = flag.String("api_version", "3.5.0", "The remote Mirth server API version")
var remote_username = flag.String("username", "", "The remote user name")
var remote_password = flag.String("password", "", "The remote password")
var verboseMode = flag.Bool("verbose", false, "Verbose console output.")

var t tracer.Tracer

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	t := tracer.New(*verboseMode)

	t.Always("Mirth API")
	mirthagent.Tracer = *t

	m := mirthagent.New(*remote_server, *remote_port)

	fmt.Println(m)
	m.Login(*remote_username, *remote_password)

}
