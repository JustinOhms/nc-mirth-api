package resource

import "fmt"

type channelGroups struct {
	p *Paths
}

func (Ω *channelGroups) BulkUpdate() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channelgroups/_bulkUpdate?override=false", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *channelGroups) GetList() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/channelgroups/", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}
