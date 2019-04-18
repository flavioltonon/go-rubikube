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

func (c *cube) Shuffle() {
	for i := 0; i < 20; i++ {
		var clockwise = rand.Int31n(1) == 1

		switch faces[rand.Int31n(5)] {
		case "front":
			c.Rotate(clockwise).Front()
			// case "back":
			// 	c.Rotate(clockwise).Back()
			// case "left":
			// 	c.Rotate(clockwise).Left()
			// case "right":
			// 	c.Rotate(clockwise).Right()
			// case "up":
			// 	c.Rotate(clockwise).Up()
			// case "down":
			// 	c.Rotate(clockwise).Down()
		}
	}
}

func (c *cube) Rotate(clockwise bool) *rotationOption {
	return &rotationOption{
		cube:      c,
		clockwise: clockwise,
	}
}

func (c *cube) Print() *printOption {
	return &printOption{
		cube: c,
	}
}

type face struct {
	name   string
	colors [][]color
}

var faces = []string{"front", "back", "left", "right", "up", "down"}

func (f face) print() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := f.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

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

type rotationOption struct {
	cube      *cube
	clockwise bool
}

func (r *rotationOption) Front() {
	var (
		front = r.cube.front.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)

	if r.clockwise {
		// Handle front rotation
		for i := 0; i < 2; i++ {
			front[0][i], front[i][2], front[2][2-i], front[2-i][0] = front[2-i][0], front[0][i], front[i][2], front[2][2-i]
		}

		// Handle sides rotation
		for i := 0; i < 3; i++ {
			tmp[i] = up[2][i]
			up[2][i] = left[2-i][2]
			left[2-i][2] = down[0][2-i]
			down[0][2-i] = right[i][0]
			right[i][0] = tmp[i]
		}
	} else {
		// Handle front rotation
		for i := 0; i < 2; i++ {
			front[0][i], front[i][2], front[2][2-i], front[2-i][0] = front[i][2-i], front[2][2-i], front[2-i][0], front[0][i]
		}

		// Handle sides rotation
		for i := 0; i < 3; i++ {
			tmp[i] = up[2][i]
			up[2][i] = right[i][0]
			right[i][0] = down[0][2-i]
			down[0][2-i] = left[2-i][2]
			left[2-i][2] = tmp[i]
		}
	}
}

func (r *rotationOption) Back() {
	var (
		back  = r.cube.back.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)

	if r.clockwise {
		// Handle back rotation
		for i := 0; i < 2; i++ {
			back[0][i], back[i][2], back[2][2-i], back[2-i][0] = back[2-i][0], back[0][i], back[i][2], back[2][2-i]
		}

		// Handle sides rotation
		for i := 0; i < 3; i++ {
			tmp[i] = up[0][2-i]
			up[0][2-i] = left[2-i][2]
			left[2-i][2] = down[0][2-i]
			down[0][2-i] = right[i][0]
			right[i][0] = tmp[i]
		}
	} else {
		// Handle back rotation
		for i := 0; i < 2; i++ {
			back[0][i], back[i][2], back[2][2-i], back[2-i][0] = back[i][2-i], back[2][2-i], back[2-i][0], back[0][i]
		}

		// Handle sides rotation
		for i := 0; i < 3; i++ {
			tmp[i] = up[0][2-i]
			up[2][i] = right[i][0]
			right[i][0] = down[0][2-i]
			down[0][2-i] = left[2-i][2]
			left[2-i][2] = tmp[i]
		}
	}
}

type printOption struct {
	cube *cube
}

func (p *printOption) Front() {
	p.cube.front.print()
}

func (p *printOption) Back() {
	p.cube.back.print()
}

func (p *printOption) Left() {
	p.cube.left.print()
}

func (p *printOption) Right() {
	p.cube.right.print()
}

func (p *printOption) Up() {
	p.cube.up.print()
}

func (p *printOption) Down() {
	p.cube.down.print()
}
