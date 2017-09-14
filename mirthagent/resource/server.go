package resource

import "fmt"

type server struct {
	p *Paths
}

func (Ω *server) GlobalScripts() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/server/globalScripts", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}
