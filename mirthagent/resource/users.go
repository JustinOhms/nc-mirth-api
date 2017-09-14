package resource

import "fmt"

type users struct {
	p *Paths
}

func (Ω *users) Login() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}

func (Ω *users) Current() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/current", Ω.p.mirthServerURL, Ω.p.mirthServerPort)
}
