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

func (Ω *channels) SetEnable() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_setEnabled", Ω.p.server, Ω.p.port)
}

func (Ω *channels) Deploy() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_deploy", Ω.p.server, Ω.p.port)
}

func (Ω *channels) Undeploy() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_undeploy", Ω.p.server, Ω.p.port)
}

func (Ω *channels) Save(channelId string) string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/%s?override=true", Ω.p.server, Ω.p.port, channelId)
}
