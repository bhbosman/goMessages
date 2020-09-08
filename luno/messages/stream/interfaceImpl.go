package stream

func (x *CreateUpdate) ReduceVolume(base float64) (leftOverVolume float64) {
	x.Volume -= base
	return x.Volume
}

func (x *Order) ReduceVolume(base float64) (leftOverVolume float64) {
	x.Volume -= base
	return x.Volume
}
