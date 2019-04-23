package rubikube

import "strings"

type faceOption struct {
	face *face
	cube *cube
}

func (f *faceOption) MainColor() {
	f.face.colors[1][1].Color().Println(strings.ToUpper(string(f.face.colors[1][1])))
}

func (f *faceOption) Position(i, j int) {
	f.face.colors[i][j].Color().Println(strings.ToUpper(string(f.face.colors[i][j])))
}

func (f *faceOption) Rotate() *rotationOption {
	return &rotationOption{
		face: f.face,
		cube: f.cube,
	}
}

func (f *faceOption) Print() *printOption {
	return &printOption{
		face: f.face,
		cube: f.cube,
	}
}
