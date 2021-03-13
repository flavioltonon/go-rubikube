package rubikube

func (c *Rubikube) Print() {
	{
		var (
			w              = white.Color()
			initialSpacing = "         "
		)

		// Build upper part of the cube
		w.Printf("\n%s__________________\n", initialSpacing)
		for i := 0; i < 6; i++ {
			spacing := initialSpacing[(3*i+2)/2:]

			w.Printf("%s/", spacing)

			for j := 0; j < 3; j++ {
				color := c.up.colors[i/2][j].Color()
				color.Print(`\\\\\`)
				w.Printf("/")
			}

			switch i % 2 {
			case 1:
				for k := 0; k < (3*i+2)/2; k++ {
					if (k-1)%3 == 0 {
						w.Printf("|")
					}
					if (k-2)%3 == 0 {
						w.Printf("/")
					}
					if (k-3)%3 == 0 {
						color := c.right.colors[k/3][k/3-i/2+2].Color()
						color.Print(`\`)
					}
				}

				w.Printf("\n")

				w.Printf("%s/-----------------/", initialSpacing[(3*i+3)/2:])
				for l := 0; l < (i+1)/2; l++ {
					color := c.right.colors[l][l-i/2+2].Color()
					color.Print(`\\`)
					w.Printf("|")
				}
				w.Printf("\n")
			case 0:
				for k := 1; k < (3*i+4)/2; k++ {
					if (k-1)%3 == 0 {
						w.Printf("|")
					}
					if (k-2)%3 == 0 {
						color := c.right.colors[k/3][k/3-i/2+3].Color()
						color.Print(`\`)
					}
					if (k-3)%3 == 0 {
						w.Printf("/")
					}
				}

				w.Printf("\n")
			}
		}

		// Build lower part of the cube
		for i := 0; i < 6; i++ {
			w.Printf("|")
			for j := 0; j < 3; j++ {
				color := c.front.colors[i/2][j].Color()
				color.Print("/////")
				w.Printf("|")
			}

			switch i % 2 {
			case 1:
				for k := 9 - 3*i/2; k > 1; k-- {
					if (k+1)%3 == 0 {
						w.Printf(`/`)
					}
					if (k-1)%3 == 0 {
						color := c.right.colors[3-k/3][2-k/3-i/2].Color()
						color.Print(`\`)
					}
					if (k-0)%3 == 0 {
						w.Printf(`|`)
					}
				}
				w.Printf("\n")

				w.Printf("|-----------------|")
				for l := 3 - (i+1)/2; l > 0; l-- {
					color := c.right.colors[3-l][2-l-i/2].Color()
					color.Print(`\\`)
					w.Printf("/")
				}
				w.Printf("\n")
			case 0:
				for k := 8 - (3*i)/2; k > 0; k-- {
					color := c.right.colors[2-k/3][2-k/3-i/2].Color()
					if (k+1)%3 == 0 {
						color.Print(`\`)
					}
					if (k-1)%3 == 0 {
						w.Printf(`/`)
					}
					if (k-0)%3 == 0 {
						w.Printf(`|`)
					}
				}

				w.Printf("\n")
			}
		}

		w.Printf("\n")
	}
}

type printHandler struct{ c *Rubikube }

func (c *Rubikube) Print2D() *printHandler { return &printHandler{c} }

func (h *printHandler) Front() { h.c.front.print2D() }
func (h *printHandler) Back()  { h.c.back.print2D() }
func (h *printHandler) Left()  { h.c.left.print2D() }
func (h *printHandler) Right() { h.c.right.print2D() }
func (h *printHandler) Up()    { h.c.up.print2D() }
func (h *printHandler) Down()  { h.c.down.print2D() }

func (f *face) print2D() {
	w := white.Color()

	w.Println(" _________________")

	for i := 0; i < 6; i++ {
		w.Printf("|")

		for j := 0; j < 3; j++ {
			color := f.colors[i/2][j].Color()
			color.Print("/////")
			w.Printf("|")
		}

		w.Printf("\n")

		if i%2 != 0 {
			w.Println("|-----------------|")
		}
	}
}
