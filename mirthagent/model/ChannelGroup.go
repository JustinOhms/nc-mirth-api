package model

import (
	"encoding/xml"
)

type ChannelGroupChannel struct {
	XMLName  xml.Name `xml:"channel"`
	Versionø string   `xml:"version,attr"`
	Idø      string   `xml:"id"`
}

func (Ω *ChannelGroupChannel) Id() string {
	return Ω.Idø
}

func (Ω *ChannelGroupChannel) SetId(v string) {
	Ω.Idø = v
}

type ChannelGroup struct {
	XMLName      xml.Name              `xml:"channelGroup"`
	Versionø     string                `xml:"version,attr"`
	Idø          string                `xml:"id"`
	Nameø        string                `xml:"name"`
	Descriptionø string                `xml:"description"`
	Channelsø    []ChannelGroupChannel `xml:"channels>channel"`
}

func NewChannelGroup(id string, name string, description string, channels ChannelIdIterator) ChannelGroup {
	c := ChannelGroup{Idø: id, Nameø: name, Descriptionø: description, Versionø: apiVersion()}
	c.AddChannels(channels)
	return c
}

//Creates a []ChannelGroup from a []ChannelGroupInterface
func NewChannelGroups(cgis []ChannelGroupInterface) []ChannelGroup {
	cg := make([]ChannelGroup, 0)

	for _, cgi := range cgis {
		ngc := NewChannelGroup(cgi.GroupId(), cgi.Name(), cgi.Description(), ChannelIdIterator(cgi))
		cg = append(cg, ngc)
	}
	return cg
}

func (Ω *ChannelGroup) Id() string {
	return Ω.Idø
}

func (Ω *ChannelGroup) SetId(v string) {
	Ω.Idø = v
}

func (Ω *ChannelGroup) Name() string {
	return Ω.Nameø
}

func (Ω *ChannelGroup) SetName(v string) {
	Ω.Nameø = v
}

func (Ω *ChannelGroup) Description() string {
	return Ω.Descriptionø
}

func (Ω *ChannelGroup) SetDescription(v string) {
	Ω.Descriptionø = v
}

func (Ω *ChannelGroup) Channels() []ChannelGroupChannel {
	return Ω.Channelsø
}

func (Ω *ChannelGroup) SetChannels(v []ChannelGroupChannel) {
	Ω.Channelsø = v
}

func (Ω *ChannelGroup) AppendChannel(v ChannelGroupChannel) {
	Ω.Channelsø = append(Ω.Channelsø, v)
}

func (Ω *ChannelGroup) AddChannel(v ChannelId) {
	c := ChannelGroupChannel{Idø: v.ChannelId(), Versionø: apiVersion()}
	Ω.AppendChannel(c)
}

func (Ω *ChannelGroup) AddChannels(chIdIt ChannelIdIterator) {
	c, done := chIdIt.ChannelIdIterator()
	for v := range c {
		Ω.AddChannel(v)
	}
	done <- true
}

type multiChannelGroup struct {
	XMLName       xml.Name       `xml:"list"`
	ChannelGroups []ChannelGroup `xml:"channelGroup"`
}

func ChannelGroupsFromXml(b []byte) []ChannelGroup {
	m := multiChannelGroup{}
	xml.Unmarshal(b, &m)
	return m.ChannelGroups
}
