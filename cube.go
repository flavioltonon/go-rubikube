package rubikube

import (
	"math/rand"

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

func (c *cube) Rotate() *rotationOption {
	return &rotationOption{
		cube: c,
	}
}

func (c *cube) Print() *printOption {
	return &printOption{
		cube: c,
	}
}

func (c *cube) Shuffle() {
	for i := 0; i < 20; i++ {
		var clockwise = rand.Int31n(1) == 1

		switch faces[rand.Int31n(5)] {
		case "front":
			if clockwise {
				c.Rotate().Clockwise().Front()
			}
			c.Rotate().CounterClockwise().Front()
			// case "back":
			// 	c.Rotate().Back()
			// case "left":
			// 	c.Rotate().Left()
			// case "right":
			// 	c.Rotate().Right()
			// case "up":
			// 	c.Rotate().Up()
			// case "down":
			// 	c.Rotate().Down()
		}
	}
}

type face struct {
	name   string
	colors [][]color
}

var faces = []string{"front", "back", "left", "right", "up", "down"}

func (f face) MainColor() color {
	return f.colors[1][1]
}

func (f face) Position(i, j int) color {
	return f.colors[i][j]
}

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
