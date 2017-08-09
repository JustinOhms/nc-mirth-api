package mirthagent

import "fmt"

func (a *Agent) loginPath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/_login", a.Server, a.Port)
}

func (a *Agent) currentUserpath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/users/current", a.Server, a.Port)
}

func (a *Agent) infoPath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/system/info", a.Server, a.Port)
}

func (a *Agent) channelStatusPath() string {
	return fmt.Sprintf("https://%s:%s/mirth/api/3.5.0/channels/statuses?includeUndeployed=true", a.Server, a.Port)
}
