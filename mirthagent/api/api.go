package api

import (
	"github.com/NavigatingCancer/mirth-api/mirthagent/api/system"
	"github.com/NavigatingCancer/mirth-api/mirthagent/session"
)

type API struct {
	Session *session.Session
	//Channels   channels
	System system.System
}

func New(s *session.Session) *API {
	api := API{Session: s}
	api.System = system.System{Session: *s}
	//	api.Channels = channels{a: *a}
	//	api.System = system{a: *a}
	return &api
}

//type channels struct {
//	a requestagent.Agent
//}

//type system struct {
//	a requestagent.Agent
//}
