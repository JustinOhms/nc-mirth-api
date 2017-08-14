package resource

import "fmt"

type channels struct {
	p *Paths
}

func (Ω *channels) Statuses() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/statuses?includeUndeployed=true", Ω.p.server, Ω.p.port)
}

func (Ω *channels) SetIntialState() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_setInitialState", Ω.p.server, Ω.p.port)
}
