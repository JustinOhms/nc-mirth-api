package model

import "encoding/xml"

type ChannelStatus struct {
	ChannelId string `xml:"channelId"`
	Name      string `xml:"name"`
	State     string `xml:"state"`
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
		m[cs.ChannelId] = cs
	}
	return m
}
