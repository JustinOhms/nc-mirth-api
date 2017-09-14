package resource

import "fmt"

type system struct {
	p *Paths
}

func (Ω *system) Info() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/system/info", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}
