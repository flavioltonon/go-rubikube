package rubikube

import (
	"fmt"

	rainbow "github.com/fatih/color"
)

type printOptionTridimensional struct {
	cube *cube
}

func (p *printOptionTridimensional) Front() {
	var (
		white          = rainbow.New(rainbow.FgWhite)
		initialSpacing = "         "
	)

	// Build upper part of the cube
	white.Print(fmt.Sprintf("%s__________________\n", initialSpacing))
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
					x, y := func(i, k int) (int, int) {
						switch i {
						case 1:
							switch k {
							case 0:
								return 0, 2
							default:
								return 0, 0
							}
						case 3:
							switch k {
							case 0:
								return 0, 1
							case 3:
								return 1, 2
							default:
								return 0, 0
							}
						case 5:
							switch k {
							case 0:
								return 0, 0
							case 3:
								return 1, 1
							case 6:
								return 2, 2
							default:
								return 0, 0
							}
						default:
							return 0, 0
						}
					}(i, k)

					color := p.cube.right.colors[x][y].Color()

					color.Print(`\`)
				}
			}

			white.Print("\n")

			white.Print(fmt.Sprintf("%s/-----------------/", initialSpacing[(3*i+3)/2:]))
			for l := 0; l < (i+1)/2; l++ {
				x, y := func(i, l int) (int, int) {
					switch i {
					case 1:
						switch l {
						case 0:
							return 0, 2
						default:
							return 0, 0
						}
					case 3:
						switch l {
						case 0:
							return 0, 1
						case 1:
							return 1, 2
						default:
							return 0, 0
						}
					case 5:
						switch l {
						case 0:
							return 0, 0
						case 1:
							return 1, 1
						case 2:
							return 2, 2
						default:
							return 0, 0
						}
					default:
						return 0, 0
					}
				}(i, l)

				color := p.cube.right.colors[x][y].Color()
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
					x, y := func(i, k int) (int, int) {
						switch i {
						case 2:
							switch k {
							case 2:
								return 0, 2
							default:
								return 0, 0
							}
						case 4:
							switch k {
							case 2:
								return 0, 1
							case 5:
								return 1, 2
							default:
								return 0, 0
							}
						default:
							return 0, 0
						}
					}(i, k)

					color := p.cube.right.colors[x][y].Color()

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
					x, y := func(i, k int) (int, int) {
						switch i {
						case 1:
							switch k {
							case 7:
								return 1, 0
							case 4:
								return 2, 1
							default:
								return 0, 0
							}
						case 3:
							switch k {
							case 4:
								return 2, 0
							default:
								return 0, 0
							}
						default:
							return 0, 0
						}
					}(i, k)

					color := p.cube.right.colors[x][y].Color()
					color.Print(`\`)
				}
				if (k-0)%3 == 0 {
					white.Print(`|`)
				}
			}

			white.Print("\n")

			white.Print("|-----------------|")
			for l := 3 - (i+1)/2; l > 0; l-- {
				x, y := func(i, l int) (int, int) {
					switch i {
					case 1:
						switch l {
						case 2:
							return 1, 0
						case 1:
							return 2, 1
						default:
							return 0, 0
						}
					case 3:
						switch l {
						case 1:
							return 2, 0
						default:
							return 0, 0
						}
					default:
						return 0, 0
					}
				}(i, l)

				color := p.cube.right.colors[x][y].Color()
				color.Print(`\\`)
				white.Print("/")
			}
			white.Print("\n")
		case 0:
			for k := 8 - (3*i)/2; k > 0; k-- {
				x, y := func(i, k int) (int, int) {
					switch i {
					case 0:
						switch k {
						case 8:
							return 0, 0
						case 5:
							return 1, 1
						case 2:
							return 2, 2
						default:
							return 0, 0
						}
					case 2:
						switch k {
						case 5:
							return 1, 0
						case 2:
							return 2, 1
						default:
							return 0, 0
						}
					case 4:
						switch k {
						case 2:
							return 2, 0
						default:
							return 0, 0
						}
					default:
						return 0, 0
					}
				}(i, k)

				color := p.cube.right.colors[x][y].Color()
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
}
