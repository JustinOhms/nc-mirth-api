package api

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/channel"
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/channelgroup"
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/codeTemplateLibraries"
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/server"
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/system"
	"github.com/NavigatingCancer/mirth-api/mirthagent/session"
)

type API struct {
	Session               *session.Session
	Channel               channel.Channel
	System                system.System
	Server                server.Server
	ChannelGroup          channelgroup.ChannelGroup
	CodeTemplateLibraries codeTemplateLibraries.CodeTemplateLibraries
}

func New(s *session.Session) *API {
	api := API{Session: s}
	api.Server = server.Server{Session: s}
	api.System = system.System{Session: s}
	api.Channel = channel.Channel{Session: s}
	api.ChannelGroup = channelgroup.ChannelGroup{Session: s}
	api.CodeTemplateLibraries = codeTemplateLibraries.CodeTemplateLibraries{Session: s}
	return &api
}
