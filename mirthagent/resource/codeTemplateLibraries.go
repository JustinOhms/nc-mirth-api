package resource

import "fmt"

type codeTemplateLibraries struct {
	p *Paths
}

func (Ω *codeTemplateLibraries) BulkUpdate() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/codeTemplateLibraries/_bulkUpdate?override=true", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}
