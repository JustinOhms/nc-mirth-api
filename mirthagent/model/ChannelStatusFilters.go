package model

type ChannelStatusFilter func(ch ChannelStatus) bool

type channelStatusFilters struct {
}

func (Ω *channelStatusFilters) State(state string) ChannelStatusFilter {
	f := func(ch ChannelStatus) bool {
		return (ch.State == state)
	}
	return f
}

func (Ω *channelStatusFilters) NotState(state string) ChannelStatusFilter {
	f := func(ch ChannelStatus) bool {
		return !(ch.State == state)
	}
	return f
}
