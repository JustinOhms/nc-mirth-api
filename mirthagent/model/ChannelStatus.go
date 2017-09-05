package model

import "encoding/xml"

type ChannelStatus struct {
	ChannelIdø string `xml:"channelId"`
	Name       string `xml:"name"`
	State      string `xml:"state"`
}

func (Ω ChannelStatus) ChannelIdƒ() string {
	return Ω.ChannelIdø
}

func (Ω ChannelStatus) ChannelId() string {
	return Ω.ChannelIdø
}

func (Ω ChannelStatus) SetChannelId(v string) {
	Ω.ChannelIdø = v
}

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
