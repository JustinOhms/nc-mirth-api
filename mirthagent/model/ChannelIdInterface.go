package model

type ChannelId interface {
	ChannelIdƒ() string
}

type ChannelIdIterator interface {
	ChannelIdIterator() (chan ChannelId, chan bool)
}
