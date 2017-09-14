package resource

import "fmt"

type channelGroups struct {
	p *Paths
}

func (Ω *channelGroups) BulkUpdate() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channelgroups/_bulkUpdate?override=false", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}
