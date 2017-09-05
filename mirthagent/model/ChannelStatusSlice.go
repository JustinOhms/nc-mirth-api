package model

type ChannelStatusSlice struct {
	Slice   []ChannelStatus
	Filters channelStatusFilters
}

func (Ω *ChannelStatusSlice) Filter(filter ChannelStatusFilter) (chan ChannelStatus, chan bool) {
	o := make(chan ChannelStatus, 1)
	ctrl := make(chan bool, 1)
	go func() {
		for _, v := range Ω.Slice {
			select {
			case done := <-ctrl:
				if done {
					close(o)
					close(ctrl)
				}
			default:
				if filter(v) {
					o <- v
				}

			}
		}
		close(o)
	}()
	return o, ctrl
}

func (Ω ChannelStatusSlice) ChannelIdProviderChannel() (chan ChannelIdProvider, chan bool) {
	o := make(chan ChannelIdProvider, 1)
	ctrl := make(chan bool, 1)
	go func() {
		for _, v := range Ω.Slice {
			select {
			case done := <-ctrl:
				if done {
					close(o)
					close(ctrl)
				}
			default:
				o <- ChannelIdProvider(v)
			}
		}
		close(o)
	}()
	return o, ctrl
}

func (Ω *ChannelStatusSlice) ToMap() map[string]ChannelStatus {
	return Ω.ToMapById()
}

func (Ω *ChannelStatusSlice) ToMapById() map[string]ChannelStatus {
	m := make(map[string]ChannelStatus)
	for _, cc := range Ω.Slice {
		m[cc.ChannelId] = cc
	}
	return m
}

func (Ω *ChannelStatusSlice) ToMapByName() map[string]ChannelStatus {
	m := make(map[string]ChannelStatus)
	for _, cc := range Ω.Slice {
		m[cc.Name] = cc
	}
	return m
}
