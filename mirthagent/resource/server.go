package resource

import (
	"fmt"
	"net/url"
)

type server struct {
	p *Paths
}

func (Ω *server) GlobalScripts() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/server/globalScripts", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *server) ResourceReload(resourceId string) string {
	escapedId := url.PathEscape(resourceId)
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/server/resources/%s/_reload", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion, escapedId)
}

func (Ω *server) ConfigurationMap() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/server/configurationMap", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *server) Resources() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/server/resources/", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}
