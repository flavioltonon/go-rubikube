package rubikube

type printOption struct {
	cube *cube
}

func (p *printOption) Bidimensional() *printOptionBidimensional {
	return &printOptionBidimensional{
		cube: p.cube,
	}
}

func (p *printOption) Tridimensional() *printOptionTridimensional {
	return &printOptionTridimensional{
		cube: p.cube,
	}
}
