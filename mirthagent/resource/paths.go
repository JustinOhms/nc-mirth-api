package resource

import "fmt"

type Paths struct {
	server   string
	port     string
	Users    users
	System   system
	Channels channels
}

func PathsNew(server string, port string) Paths {
	p := Paths{server: server, port: port}
	p.Users = users{p: &p}
	p.System = system{p: &p}
	p.Channels = channels{p: &p}
	return p
}

type users struct {
	p *Paths
}

func (Ω *users) Login() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", Ω.p.server, Ω.p.port)
}

func (Ω *users) Current() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/current", Ω.p.server, Ω.p.port)
}

type system struct {
	p *Paths
}

func (Ω *system) Info() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/system/info", Ω.p.server, Ω.p.port)
}

type channels struct {
	p *Paths
}

func (Ω *channels) Statuses() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/statuses?includeUndeployed=true", Ω.p.server, Ω.p.port)
}

func (Ω *channels) SetIntialState() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_setInitialState", Ω.p.server, Ω.p.port)
}
