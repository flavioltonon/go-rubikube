package rubikube

import (
	"math/rand"
	"time"

	rainbow "github.com/fatih/color"
)

type cube struct {
	front face
	back  face
	left  face
	right face
	up    face
	down  face
}

func NewCube() *cube {
	return &cube{
		front: face{colors: [][]color{{blue, blue, blue}, {blue, blue, blue}, {blue, blue, blue}}},
		back:  face{colors: [][]color{{red, red, red}, {red, red, red}, {red, red, red}}},
		left:  face{colors: [][]color{{magenta, magenta, magenta}, {magenta, magenta, magenta}, {magenta, magenta, magenta}}},
		right: face{colors: [][]color{{green, green, green}, {green, green, green}, {green, green, green}}},
		up:    face{colors: [][]color{{white, white, white}, {white, white, white}, {white, white, white}}},
		down:  face{colors: [][]color{{yellow, yellow, yellow}, {yellow, yellow, yellow}, {yellow, yellow, yellow}}},
	}
}

func (c *cube) Front() *faceOption {
	return &faceOption{
		face: &face{
			name:   "front",
			colors: c.front.colors,
		},
		cube: c,
	}
}

func (c *cube) Back() *faceOption {
	return &faceOption{
		face: &face{
			name:   "back",
			colors: c.back.colors,
		},
		cube: c,
	}
}

func (c *cube) Left() *faceOption {
	return &faceOption{
		face: &face{
			name:   "left",
			colors: c.left.colors,
		},
		cube: c,
	}
}

func (c *cube) Right() *faceOption {
	return &faceOption{
		face: &face{
			name:   "right",
			colors: c.right.colors,
		},
		cube: c,
	}
}

func (c *cube) Up() *faceOption {
	return &faceOption{
		face: &face{
			name:   "up",
			colors: c.up.colors,
		},
		cube: c,
	}
}

func (c *cube) Down() *faceOption {
	return &faceOption{
		face: &face{
			name:   "down",
			colors: c.down.colors,
		},
		cube: c,
	}
}

func (c *cube) Shuffle(n int) {
	var source = rand.NewSource(time.Now().Unix())

	r := rand.New(source)

	for i := 0; i < n; i++ {
		var face *faceOption

		switch faces[r.Int31n(6)] {
		case "front":
			face = c.Front()
		case "back":
			face = c.Back()
		case "left":
			face = c.Left()
		case "right":
			face = c.Right()
		case "up":
			face = c.Up()
		case "down":
			face = c.Down()
		}

		rotate := face.Rotate()

		if r.Int31n(100) > 50 {
			rotate.Clockwise()
		} else {
			rotate.CounterClockwise()
		}
	}
}

type face struct {
	name   string
	colors [][]color
}

var faces = []string{"front", "back", "left", "right", "up", "down"}

type color string

const (
	blue    color = "blue"
	red     color = "red"
	magenta color = "magenta"
	green   color = "green"
	white   color = "white"
	yellow  color = "yellow"
)

func (c color) Color() *rainbow.Color {
	switch c {
	case blue:
		return rainbow.New(rainbow.FgBlue)
	case red:
		return rainbow.New(rainbow.FgRed)
	case magenta:
		return rainbow.New(rainbow.FgMagenta)
	case green:
		return rainbow.New(rainbow.FgGreen)
	case white:
		return rainbow.New(rainbow.FgWhite)
	case yellow:
		return rainbow.New(rainbow.FgYellow)
	default:
		return rainbow.New(rainbow.FgBlack)
	}
}
