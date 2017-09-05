package model

type ChannelId interface {
	ChannelId() string
}

type ChannelIdIterator interface {
	ChannelIdIterator() (chan ChannelId, chan bool)
}
