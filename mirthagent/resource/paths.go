package resource

type Paths struct {
	server        string
	port          string
	Users         users
	System        system
	Channels      channels
	ChannelGroups channelGroups
}

func PathsNew(server string, port string) Paths {
	p := Paths{server: server, port: port}
	p.Users = users{p: &p}
	p.System = system{p: &p}
	p.Channels = channels{p: &p}
	p.ChannelGroups = channelGroups{p: &p}
	return p
}
