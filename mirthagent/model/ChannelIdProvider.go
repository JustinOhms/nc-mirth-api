package model

type ChannelIdProvider interface {
	ChannelIdƒ() string
}

type ChannelIdProviderChannelProvider interface {
	ChannelIdProviderChannel() (chan ChannelIdProvider, chan bool)
}
