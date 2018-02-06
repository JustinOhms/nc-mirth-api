package resource

type Paths struct {
	mirthServerURL        string
	mirthServerPort       string
	mirthServerVersion    string
	Users                 users
	Server                server
	System                system
	Channels              channels
	ChannelGroups         channelGroups
	CodeTemplateLibraries codeTemplateLibraries
}

func PathsNew(serverURL string, serverPort string, serverVersion string) Paths {
	p := Paths{mirthServerURL: serverURL, mirthServerPort: serverPort, mirthServerVersion: serverVersion}
	p.Users = users{p: &p}
	p.Server = server{p: &p}
	p.System = system{p: &p}
	p.Channels = channels{p: &p}
	p.ChannelGroups = channelGroups{p: &p}
	p.CodeTemplateLibraries = codeTemplateLibraries{p: &p}
	return p
}
