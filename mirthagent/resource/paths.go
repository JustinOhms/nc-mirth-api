package resource

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
