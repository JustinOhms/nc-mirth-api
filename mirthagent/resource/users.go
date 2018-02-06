package resource

import "fmt"

type users struct {
	p *Paths
}

func (Ω *users) Login() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/users/_login", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}

func (Ω *users) Current() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/%s/users/current", Ω.p.mirthServerURL, Ω.p.mirthServerPort, Ω.p.mirthServerVersion)
}
