package rubikube

import (
	"fmt"
	"strings"

	rainbow "github.com/fatih/color"
)

type printOption struct {
	face *face
	cube *cube
}

func (p *printOption) Bidimensional() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Println(strings.ToUpper(p.face.name) + ":")
	white.Println(" _________________")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.face.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Println("|-----------------|")
		}
	}
}

func (p *printOption) Tridimensional() {
	var (
		white          = rainbow.New(rainbow.FgWhite)
		initialSpacing = "         "
	)

	// Build upper part of the cube
	white.Println(fmt.Sprintf("\n%s__________________", initialSpacing))
	for i := 0; i < 6; i++ {
		spacing := initialSpacing[(3*i+2)/2:]
		white.Print(fmt.Sprintf("%s/", spacing))
		for j := 0; j < 3; j++ {
			color := p.cube.up.colors[i/2][j].Color()
			color.Print(`\\\\\`)
			white.Print("/")
		}

		switch i % 2 {
		case 1:
			for k := 0; k < (3*i+2)/2; k++ {
				if (k-1)%3 == 0 {
					white.Print("|")
				}
				if (k-2)%3 == 0 {
					white.Print(`/`)
				}
				if (k-3)%3 == 0 {
					color := p.cube.right.colors[k/3][k/3-i/2+2].Color()
					color.Print(`\`)
				}
			}

			white.Print("\n")

			white.Print(fmt.Sprintf("%s/-----------------/", initialSpacing[(3*i+3)/2:]))
			for l := 0; l < (i+1)/2; l++ {
				color := p.cube.right.colors[l][l-i/2+2].Color()
				color.Print(`\\`)
				white.Print("|")
			}
			white.Print("\n")
		case 0:
			for k := 1; k < (3*i+4)/2; k++ {
				if (k-1)%3 == 0 {
					white.Print("|")
				}
				if (k-2)%3 == 0 {
					color := p.cube.right.colors[k/3][k/3-i/2+3].Color()
					color.Print(`\`)
				}
				if (k-3)%3 == 0 {
					white.Print(`/`)
				}
			}

			white.Print("\n")
		}
	}

	// Build lower part of the cube
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.front.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}

		switch i % 2 {
		case 1:
			for k := 9 - 3*i/2; k > 1; k-- {
				if (k+1)%3 == 0 {
					white.Print(`/`)
				}
				if (k-1)%3 == 0 {
					color := p.cube.right.colors[3-k/3][2-k/3-i/2].Color()
					color.Print(`\`)
				}
				if (k-0)%3 == 0 {
					white.Print(`|`)
				}
			}

			white.Print("\n")

			white.Print("|-----------------|")
			for l := 3 - (i+1)/2; l > 0; l-- {
				color := p.cube.right.colors[3-l][2-l-i/2].Color()
				color.Print(`\\`)
				white.Print("/")
			}
			white.Print("\n")
		case 0:
			for k := 8 - (3*i)/2; k > 0; k-- {
				color := p.cube.right.colors[2-k/3][2-k/3-i/2].Color()
				if (k+1)%3 == 0 {
					color.Print(`\`)
				}
				if (k-1)%3 == 0 {
					white.Print(`/`)
				}
				if (k-0)%3 == 0 {
					white.Print(`|`)
				}
			}

			white.Print("\n")
		}
	}

	white.Print("\n")
}
