package model

import "encoding/xml"

type ChannelStatus struct {
	ChannelIdø string `xml:"channelId"`
	Nameø      string `xml:"name"`
	Stateø     string `xml:"state"`
}

// ------------------------------------------------

func (Ω ChannelStatus) ChannelId() string {
	return Ω.ChannelIdø
}

func (Ω ChannelStatus) SetChannelId(v string) {
	Ω.ChannelIdø = v
}

func (Ω ChannelStatus) Name() string {
	return Ω.Nameø
}

func (Ω ChannelStatus) SetName(v string) {
	Ω.Nameø = v
}

func (Ω ChannelStatus) State() string {
	return Ω.Stateø
}

func (Ω ChannelStatus) SetState(v string) {
	Ω.Stateø = v
}

// ------------------------------------------------

type multiChannelStatus struct {
	XMLName  xml.Name        `xml:"list"`
	Channels []ChannelStatus `xml:"dashboardStatus"`
}

func ChannelStatusFromXml(b []byte) []ChannelStatus {
	m := multiChannelStatus{}
	xml.Unmarshal(b, &m)
	return m.Channels
}

func ChannelStatusArrayToMap(a []ChannelStatus) map[string]ChannelStatus {
	m := make(map[string]ChannelStatus)
	for _, cs := range a {
		m[cs.ChannelId()] = cs
	}
	return m
}
