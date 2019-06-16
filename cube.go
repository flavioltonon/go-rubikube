package rubikube

import (
	"fmt"
	"strings"

	rainbow "github.com/fatih/color"
)

var faces = make([]string, 0)

func init() {
	for _, face := range standardFaces {
		faces = append(faces, face.name)
	}
}

type cube struct {
	faces map[string]face
}

func NewCube() *cube {
	var cube cube

	faces := make(map[string]face, 0)
	for _, f := range standardFaces {
		switch f.name {
		case FACE_FRONT_NAME:
			f.colors = &[][]color{
				{blue, blue, blue},
				{blue, blue, blue},
				{blue, blue, blue},
			}
		case FACE_BACK_NAME:
			f.colors = &[][]color{
				{red, red, red},
				{red, red, red},
				{red, red, red},
			}
		case FACE_LEFT_NAME:
			f.colors = &[][]color{
				{magenta, magenta, magenta},
				{magenta, magenta, magenta},
				{magenta, magenta, magenta},
			}
		case FACE_RIGHT_NAME:
			f.colors = &[][]color{
				{green, green, green},
				{green, green, green},
				{green, green, green},
			}
		case FACE_UP_NAME:
			f.colors = &[][]color{
				{white, white, white},
				{white, white, white},
				{white, white, white},
			}
		case FACE_DOWN_NAME:
			f.colors = &[][]color{
				{yellow, yellow, yellow},
				{yellow, yellow, yellow},
				{yellow, yellow, yellow},
			}
		}

		faces[f.name] = f
	}

	cube.faces = faces

	return &cube
}

func parseCube(snapshot snapshot) *cube {
	var c cube

	c.faces = make(map[string]face, 0)

	for _, faceString := range strings.Split(string(snapshot), ";") {
		if len(faceString) == 0 {
			break
		}

		faceName := strings.Split(faceString, ":")[0]

		faceColors := make([][]color, 0)
		lineColors := make([]color, 0)
		for index, colorName := range strings.Split(strings.Split(faceString, ":")[1], ",") {
			line := index / 3
			column := index % 3

			if line == 0 {
				lineColors = make([]color, 3)
			}

			lineColors[column], _ = parseColor(colorName)

			if column == 2 {
				faceColors = append(faceColors, lineColors)

				if line == 2 {
					c.faces[faceName] = face{
						name:   faceName,
						colors: &faceColors,
					}
				}
			}

		}
	}

	return &c
}

func (c *cube) Shuffle(n int) {
	for i := 0; i < n; i++ {
		c.apply(randomMove())
	}
}

func (c *cube) Copy() *cube {
	// copy faces
	faces := make(map[string]face)
	for k, v := range c.faces {
		var colors = make([][]color, 0)
		for _, line := range v.Colors() {
			newLine := make([]color, len(line))
			copy(newLine, line)
			colors = append(colors, newLine)
		}
		faces[k] = face{
			name:   k,
			colors: &colors,
		}
	}

	return &cube{
		faces: faces,
	}
}

func (c *cube) Solve() *solveOption {
	return &solveOption{
		c,
	}
}

func (c *cube) apply(moves ...move) {
	for _, m := range moves {
		switch m.instruction() {
		case INSTRUCTION_ROTATE_FRONT_CLOCKWISE:
			c.Front().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_FRONT_COUNTERCLOCKWISE:
			c.Front().Rotate().CounterClockwise()
		case INSTRUCTION_ROTATE_BACK_CLOCKWISE:
			c.Back().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_BACK_COUNTERCLOCKWISE:
			c.Back().Rotate().CounterClockwise()
		case INSTRUCTION_ROTATE_LEFT_CLOCKWISE:
			c.Left().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_LEFT_COUNTERCLOCKWISE:
			c.Left().Rotate().CounterClockwise()
		case INSTRUCTION_ROTATE_RIGHT_CLOCKWISE:
			c.Right().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_RIGHT_COUNTERCLOCKWISE:
			c.Right().Rotate().CounterClockwise()
		case INSTRUCTION_ROTATE_UP_CLOCKWISE:
			c.Up().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_UP_COUNTERCLOCKWISE:
			c.Up().Rotate().CounterClockwise()
		case INSTRUCTION_ROTATE_DOWN_CLOCKWISE:
			c.Down().Rotate().Clockwise()
		case INSTRUCTION_ROTATE_DOWN_COUNTERCLOCKWISE:
			c.Down().Rotate().CounterClockwise()
		}
	}
}

func (c cube) IsSolved() bool {
	for _, face := range c.faces {
		mainColor := face.Position(1, 1)
		for _, line := range face.Colors() {
			for _, color := range line {
				if color != mainColor {
					return false
				}
			}
		}
	}
	return true
}

func (c cube) Front() *faceOption {
	var front = c.faces[FACE_FRONT_NAME]
	return &faceOption{
		face: &front,
		cube: &c,
	}
}

func (c cube) Back() *faceOption {
	var back = c.faces[FACE_BACK_NAME]
	return &faceOption{
		&back,
		&c,
	}
}

func (c cube) Left() *faceOption {
	var left = c.faces[FACE_LEFT_NAME]
	return &faceOption{
		&left,
		&c,
	}
}

func (c cube) Right() *faceOption {
	var right = c.faces[FACE_RIGHT_NAME]
	return &faceOption{
		&right,
		&c,
	}
}

func (c cube) Up() *faceOption {
	var up = c.faces[FACE_UP_NAME]
	return &faceOption{
		&up,
		&c,
	}
}

func (c cube) Down() *faceOption {
	var down = c.faces[FACE_DOWN_NAME]
	return &faceOption{
		&down,
		&c,
	}
}

type snapshot string

func (s snapshot) String() string {
	return string(s)
}

var defaultSnapshot = snapshot(FACE_FRONT_NAME + ":b,b,b,b,b,b,b,b,b;" +
	FACE_BACK_NAME + ":r,r,r,r,r,r,r,r,r;" +
	FACE_LEFT_NAME + ":m,m,m,m,m,m,m,m,m;" +
	FACE_RIGHT_NAME + ":g,g,g,g,g,g,g,g,g;" +
	FACE_UP_NAME + ":w,w,w,w,w,w,w,w,w;" +
	FACE_DOWN_NAME + ":y,y,y,y,y,y,y,y,y")

func (c cube) snapshot() snapshot {
	var colors = make(map[string][]string, 0)

	for _, face := range c.faces {
		for _, line := range face.Colors() {
			for _, color := range line {
				colors[face.name] = append(colors[face.name], string(color.String()[0]))
			}
		}
	}

	return snapshot(FACE_FRONT_NAME + ":" + strings.Join(colors[FACE_FRONT_NAME], ",") + ";" +
		FACE_BACK_NAME + ":" + strings.Join(colors[FACE_BACK_NAME], ",") + ";" +
		FACE_LEFT_NAME + ":" + strings.Join(colors[FACE_LEFT_NAME], ",") + ";" +
		FACE_RIGHT_NAME + ":" + strings.Join(colors[FACE_RIGHT_NAME], ",") + ";" +
		FACE_UP_NAME + ":" + strings.Join(colors[FACE_UP_NAME], ",") + ";" +
		FACE_DOWN_NAME + ":" + strings.Join(colors[FACE_DOWN_NAME], ","))
}

type face struct {
	name   string
	colors *[][]color
}

var standardFaces = []face{
	{name: FACE_FRONT_NAME},
	{name: FACE_BACK_NAME},
	{name: FACE_LEFT_NAME},
	{name: FACE_RIGHT_NAME},
	{name: FACE_UP_NAME},
	{name: FACE_DOWN_NAME},
}

func (f face) Colors() [][]color {
	return *f.colors
}

func (f face) Position(i, j int) color {
	return f.Colors()[i][j]
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

func parseColor(name string) (color, error) {
	switch name {
	case "b", "blue":
		return blue, nil
	case "r", "red":
		return red, nil
	case "m", "magenta":
		return magenta, nil
	case "g", "green":
		return green, nil
	case "w", "white":
		return white, nil
	case "y", "yellow":
		return yellow, nil
	default:
		return "", fmt.Errorf("color not registered: %s", name)
	}
}

func (c color) String() string {
	return string(c)
}

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

type change struct {
	Index       int         `json:"index" bson:"index"`
	Instruction instruction `json:"instruction" bson:"instruction"`
}

type changes []change
