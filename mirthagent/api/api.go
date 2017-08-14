package api

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/channel"
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/system"
	"github.com/NavigatingCancer/mirth-api/mirthagent/session"
)

type API struct {
	Session *session.Session
	Channel channel.Channel
	System  system.System
}

func New(s *session.Session) *API {
	api := API{Session: s}
	api.System = system.System{Session: *s}
	api.Channel = channel.Channel{Session: *s}
	return &api
}
