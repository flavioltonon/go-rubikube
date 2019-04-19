package rubikube

type rotationOption struct {
	cube *cube
}

func (r *rotationOption) Clockwise() *rotationOptionClockwise {
	return &rotationOptionClockwise{
		cube: r.cube,
	}
}

func (r *rotationOption) CounterClockwise() *rotationOptionCounterClockwise {
	return &rotationOptionCounterClockwise{
		cube: r.cube,
	}
}
