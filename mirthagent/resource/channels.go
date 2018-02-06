package resource

import "fmt"

type channels struct {
	p *Paths
}

func (Ω *channels) Statuses() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/statuses?includeUndeployed=true", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channels) SetIntialState() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/_setInitialState", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channels) SetEnable() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/_setEnabled", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channels) Deploy() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/_deploy", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channels) Undeploy() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/_undeploy", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channels) Save(channelId string) string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channels/%s?override=true", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion, channelId)
}
