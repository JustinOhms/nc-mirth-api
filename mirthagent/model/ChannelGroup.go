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

func (Ω *ChannelGroup) Id() string {
	return Ω.Id()
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
	c := ChannelGroupChannel{Idø: v.ChannelId(), Versionø: "3.5.0"}
	Ω.AppendChannel(c)
}
