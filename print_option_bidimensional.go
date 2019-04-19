package rubikube

import rainbow "github.com/fatih/color"

type printOptionBidimensional struct {
	cube *cube
}

func (p *printOptionBidimensional) Front() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.front.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

func (p *printOptionBidimensional) Back() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.back.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

func (p *printOptionBidimensional) Left() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.left.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

func (p *printOptionBidimensional) Right() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.right.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

func (p *printOptionBidimensional) Up() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.up.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}

func (p *printOptionBidimensional) Down() {
	var white = rainbow.New(rainbow.FgWhite)

	white.Print(" _________________\n")
	for i := 0; i < 6; i++ {
		white.Print("|")
		for j := 0; j < 3; j++ {
			color := p.cube.down.colors[i/2][j].Color()
			color.Print("/////")
			white.Print("|")
		}
		white.Print("\n")
		if i%2 != 0 {
			white.Print("|-----------------|\n")
		}
	}
}
