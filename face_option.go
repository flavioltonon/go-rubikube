package rubikube

type faceOption struct {
	*face
	*cube
}

func (f *faceOption) Rotate() *rotationOption {
	return &rotationOption{
		f.face,
		f.cube,
	}
}

func (f *faceOption) Print() *printOption {
	return &printOption{
		f.face,
		f.cube,
	}
}
