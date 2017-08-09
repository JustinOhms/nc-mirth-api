package channelstatus

import (
	"encoding/xml"
)

type ChannelStatus struct {
	ChannelId string `xml:"channelId"`
	Name      string `xml:"name"`
	State     string `xml:"state"`
}

type multiChannelStatus struct {
	XMLName  xml.Name        `xml:"list"`
	Channels []ChannelStatus `xml:"dashboardStatus"`
}

func XmlParse(b []byte) []ChannelStatus {
	m := multiChannelStatus{}
	xml.Unmarshal(b, &m)
	return m.Channels
}

type Handler func([]ChannelStatus)
