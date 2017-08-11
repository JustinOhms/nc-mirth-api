package resource

import "fmt"

type Paths struct {
	server   string
	port     string
	Users    users
	System   system
	Channels channels
}

func New(server string, port string) Paths {
	p := Paths{server: server, port: port}
	p.Users = users{p: &p}
	p.System = system{p: &p}
	p.Channels = channels{p: &p}
	return p
}

type users struct {
	p *Paths
}

func (u *users) Login() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", u.p.server, u.p.port)
}

func (u *users) Current() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/current", u.p.server, u.p.port)
}

type system struct {
	p *Paths
}

func (s *system) Info() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/system/info", s.p.server, s.p.port)
}

type channels struct {
	p *Paths
}

func (c *channels) Statuses() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/statuses?includeUndeployed=true", c.p.server, c.p.port)
}

func (c *channels) SetIntialState() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/_setInitialState", c.p.server, c.p.port)
}
