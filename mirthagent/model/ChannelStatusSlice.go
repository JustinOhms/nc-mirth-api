package model

type ChannelStatusSlice struct {
	Slice          []ChannelStatus
	Filters        channelStatusFilters
	currentFilters []ChannelStatusFilter
}

//adds a filter to be used by iterators
func (Ω *ChannelStatusSlice) Filter(filter ChannelStatusFilter) *ChannelStatusSlice {
	Ω.currentFilters = append(Ω.currentFilters, filter)
	return Ω
}

//clears all filters used in iterators
func (Ω *ChannelStatusSlice) ClearFilters() {
	Ω.currentFilters = []ChannelStatusFilter{}
}

func shouldFilterOut(v ChannelStatus, filters []ChannelStatusFilter) bool {
	for _, f := range filters {
		if !f(v) {
			return true
		}
	}
	return false
}

func (Ω *ChannelStatusSlice) Iterator() (chan ChannelStatus, chan bool) {
	o := make(chan ChannelStatus, 1)
	ctrl := make(chan bool, 1)
	go func(filters []ChannelStatusFilter) {
		for _, v := range Ω.Slice {
			select {
			case done := <-ctrl:
				if done {
					close(o)
					close(ctrl)
				}
			default:
				if !shouldFilterOut(v, filters) {
					o <- v
				}

			}
		}
		close(o)
	}(Ω.currentFilters)
	return o, ctrl
}

func (Ω ChannelStatusSlice) ChannelIdIterator() (chan ChannelId, chan bool) {
	o := make(chan ChannelId, 1)
	ctrl := make(chan bool, 1)
	go func(filters []ChannelStatusFilter) {
		for _, v := range Ω.Slice {
			select {
			case done := <-ctrl:
				if done {
					close(o)
					close(ctrl)
				}
			default:
				if !shouldFilterOut(v, filters) {
					o <- ChannelId(v)
				}
			}
		}
		close(o)
	}(Ω.currentFilters)
	return o, ctrl
}

func (Ω *ChannelStatusSlice) ToMap() map[string]ChannelStatus {
	return Ω.ToMapById()
}

func (Ω *ChannelStatusSlice) ToMapById() map[string]ChannelStatus {
	m := make(map[string]ChannelStatus)
	for _, cc := range Ω.Slice {
		m[cc.ChannelId()] = cc
	}
	return m
}

func (Ω *ChannelStatusSlice) ToMapByName() map[string]ChannelStatus {
	m := make(map[string]ChannelStatus)
	for _, cc := range Ω.Slice {
		m[cc.Nameø] = cc
	}
	return m
}
