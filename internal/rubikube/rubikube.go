package rubikube

import (
	"fmt"
)

// Rubikube is a Rubik cube
type Rubikube struct {
	front face
	back  face
	left  face
	right face
	up    face
	down  face
}

// New creates a new Rubikube
func New() *Rubikube {
	return &Rubikube{
		front: DefaultFrontFace,
		back:  DefaultBackFace,
		left:  DefaultLeftFace,
		right: DefaultRightFace,
		up:    DefaultUpFace,
		down:  DefaultDownFace,
	}
}

type face struct {
	colors [3][3]color
}

var (
	DefaultFrontFace = face{
		colors: [3][3]color{
			{blue, blue, blue},
			{blue, blue, blue},
			{blue, blue, blue},
		},
	}

	DefaultBackFace = face{
		colors: [3][3]color{
			{red, red, red},
			{red, red, red},
			{red, red, red},
		},
	}

	DefaultLeftFace = face{
		colors: [3][3]color{
			{magenta, magenta, magenta},
			{magenta, magenta, magenta},
			{magenta, magenta, magenta},
		},
	}

	DefaultRightFace = face{
		colors: [3][3]color{
			{green, green, green},
			{green, green, green},
			{green, green, green},
		},
	}

	DefaultUpFace = face{
		colors: [3][3]color{
			{white, white, white},
			{white, white, white},
			{white, white, white},
		},
	}

	DefaultDownFace = face{
		colors: [3][3]color{
			{yellow, yellow, yellow},
			{yellow, yellow, yellow},
			{yellow, yellow, yellow},
		},
	}
)

func (c *Rubikube) copy() *Rubikube {
	return &Rubikube{
		front: c.front.copy(),
		back:  c.back.copy(),
		left:  c.left.copy(),
		right: c.right.copy(),
		up:    c.up.copy(),
		down:  c.down.copy(),
	}
}

func (f face) copy() face {
	var new face

	for i := range f.colors {
		for j := range f.colors[i] {
			new.colors[i][j] = f.colors[i][j]
		}
	}

	return new
}

func (c *Rubikube) do(input movement) {
	switch input {
	case ROTATE_FRONT_CLOCKWISE:
		c.Rotate().Front().Clockwise()
	case ROTATE_FRONT_COUNTERCLOCKWISE:
		c.Rotate().Front().CounterClockwise()
	case ROTATE_BACK_CLOCKWISE:
		c.Rotate().Back().Clockwise()
	case ROTATE_BACK_COUNTERCLOCKWISE:
		c.Rotate().Back().CounterClockwise()
	case ROTATE_LEFT_CLOCKWISE:
		c.Rotate().Left().Clockwise()
	case ROTATE_LEFT_COUNTERCLOCKWISE:
		c.Rotate().Left().CounterClockwise()
	case ROTATE_RIGHT_CLOCKWISE:
		c.Rotate().Right().Clockwise()
	case ROTATE_RIGHT_COUNTERCLOCKWISE:
		c.Rotate().Right().CounterClockwise()
	case ROTATE_UP_CLOCKWISE:
		c.Rotate().Up().Clockwise()
	case ROTATE_UP_COUNTERCLOCKWISE:
		c.Rotate().Up().CounterClockwise()
	case ROTATE_DOWN_CLOCKWISE:
		c.Rotate().Down().Clockwise()
	case ROTATE_DOWN_COUNTERCLOCKWISE:
		c.Rotate().Down().CounterClockwise()
	}
}

func (c *Rubikube) snapshot() string {
	return fmt.Sprintf("%s%s%s%s%s%s",
		c.front.snapshot(),
		c.back.snapshot(),
		c.left.snapshot(),
		c.right.snapshot(),
		c.up.snapshot(),
		c.down.snapshot(),
	)
}

func (f *face) snapshot() string {
	var ss string

	for i := range f.colors {
		for j := range f.colors[i] {
			ss = fmt.Sprintf("%s%d", ss, f.colors[i][j])
		}
	}

	return ss
}

func (c *Rubikube) ApplySolution(s *Solution) {
	if s == nil {
		return
	}

	for _, m := range s.movements {
		c.do(m)
	}
}
