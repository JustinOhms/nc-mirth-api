package resource

import (
	"fmt"
	"net/url"
)

type server struct {
	p *Paths
}

func (Ω *server) GlobalScripts() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/server/globalScripts", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}

func (Ω *server) ResourceReload(resourceId string) string {
	escapedId := url.PathEscape(resourceId)
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/server/resources/%s/_reload", Ω.p.mirthServerURL, Ω.p.mirthServerPort, escapedId)
}

func (Ω *server) ConfigurationMap() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/server/configurationMap", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}
