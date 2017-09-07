package model

type ChannelId interface {
	ChannelId() string
}

type ChannelIdIterator interface {
	ChannelIdIterator() (chan ChannelId, chan bool)
}

type ChannelGroupInterface interface {
	ChannelIdIterator
	Name() string
	GroupId() string
	Description() string
}
