package resource

import "fmt"

type codeTemplateLibraries struct {
	p *Paths
}

func (Ω *codeTemplateLibraries) BulkUpdate() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/codeTemplateLibraries/_bulkUpdate?override=true", Ω.p.server, Ω.p.port)
}
