package rubikube

type rotationHandler struct{ c *Rubikube }

func (c *Rubikube) Rotate() *rotationHandler { return &rotationHandler{c} }

type frontRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Front() *frontRotationHandler { return &frontRotationHandler{h.c} }

func (h *frontRotationHandler) Clockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.front.colors[0][i], h.c.front.colors[2-i][0], h.c.front.colors[2][2-i], h.c.front.colors[i][2] =
			h.c.front.colors[2-i][0], h.c.front.colors[2][2-i], h.c.front.colors[i][2], h.c.front.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[2][i]
		h.c.up.colors[2][i] = h.c.left.colors[2-i][2]
		h.c.left.colors[2-i][2] = h.c.down.colors[0][2-i]
		h.c.down.colors[0][2-i] = h.c.right.colors[i][0]
		h.c.right.colors[i][0] = tmp[i]
	}

	return h.c
}

func (h *frontRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.front.colors[0][i], h.c.front.colors[i][2], h.c.front.colors[2][2-i], h.c.front.colors[2-i][0] =
			h.c.front.colors[i][2], h.c.front.colors[2][2-i], h.c.front.colors[2-i][0], h.c.front.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[2][i]
		h.c.up.colors[2][i] = h.c.right.colors[i][0]
		h.c.right.colors[i][0] = h.c.down.colors[0][2-i]
		h.c.down.colors[0][2-i] = h.c.left.colors[2-i][2]
		h.c.left.colors[2-i][2] = tmp[i]
	}

	return h.c
}

type backRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Back() *backRotationHandler { return &backRotationHandler{h.c} }

func (h *backRotationHandler) Clockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.back.colors[0][i], h.c.back.colors[2-i][0], h.c.back.colors[2][2-i], h.c.back.colors[i][2] =
			h.c.back.colors[2-i][0], h.c.back.colors[2][2-i], h.c.back.colors[i][2], h.c.back.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[0][i]
		h.c.up.colors[0][i] = h.c.right.colors[i][2]
		h.c.right.colors[i][2] = h.c.down.colors[2][2-i]
		h.c.down.colors[2][2-i] = h.c.left.colors[2-i][0]
		h.c.left.colors[2-i][0] = tmp[i]
	}

	return h.c
}

func (h *backRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.back.colors[0][i], h.c.back.colors[i][2], h.c.back.colors[2][2-i], h.c.back.colors[2-i][0] =
			h.c.back.colors[i][2], h.c.back.colors[2][2-i], h.c.back.colors[2-i][0], h.c.back.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[0][i]
		h.c.up.colors[0][i] = h.c.left.colors[2-i][0]
		h.c.left.colors[2-i][0] = h.c.down.colors[2][2-i]
		h.c.down.colors[2][2-i] = h.c.right.colors[i][2]
		h.c.right.colors[i][2] = tmp[i]
	}

	return h.c
}

type leftRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Left() *leftRotationHandler { return &leftRotationHandler{h.c} }

func (h *leftRotationHandler) Clockwise() *Rubikube {
	// Handle left rotation
	for i := 0; i < 2; i++ {
		h.c.left.colors[0][i], h.c.left.colors[2-i][0], h.c.left.colors[2][2-i], h.c.left.colors[i][2] =
			h.c.left.colors[2-i][0], h.c.left.colors[2][2-i], h.c.left.colors[i][2], h.c.left.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[i][0]
		h.c.up.colors[i][0] = h.c.back.colors[2-i][2]
		h.c.back.colors[2-i][2] = h.c.down.colors[i][0]
		h.c.down.colors[i][0] = h.c.front.colors[i][0]
		h.c.front.colors[i][0] = tmp[i]
	}

	return h.c
}

func (h *leftRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.left.colors[0][i], h.c.left.colors[i][2], h.c.left.colors[2][2-i], h.c.left.colors[2-i][0] =
			h.c.left.colors[i][2], h.c.left.colors[2][2-i], h.c.left.colors[2-i][0], h.c.left.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[i][0]
		h.c.up.colors[i][0] = h.c.front.colors[i][0]
		h.c.front.colors[i][0] = h.c.down.colors[i][0]
		h.c.down.colors[i][0] = h.c.back.colors[2-i][2]
		h.c.back.colors[2-i][2] = tmp[i]
	}

	return h.c
}

type rightRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Right() *rightRotationHandler { return &rightRotationHandler{h.c} }

func (h *rightRotationHandler) Clockwise() *Rubikube {
	// Handle right rotation
	for i := 0; i < 2; i++ {
		h.c.right.colors[0][i], h.c.right.colors[2-i][0], h.c.right.colors[2][2-i], h.c.right.colors[i][2] =
			h.c.right.colors[2-i][0], h.c.right.colors[2][2-i], h.c.right.colors[i][2], h.c.right.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[i][2]
		h.c.up.colors[i][2] = h.c.front.colors[i][2]
		h.c.front.colors[i][2] = h.c.down.colors[i][2]
		h.c.down.colors[i][2] = h.c.back.colors[2-i][0]
		h.c.back.colors[2-i][0] = tmp[i]
	}

	return h.c
}

func (h *rightRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.right.colors[0][i], h.c.right.colors[i][2], h.c.right.colors[2][2-i], h.c.right.colors[2-i][0] =
			h.c.right.colors[i][2], h.c.right.colors[2][2-i], h.c.right.colors[2-i][0], h.c.right.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.up.colors[i][2]
		h.c.up.colors[i][2] = h.c.back.colors[2-i][0]
		h.c.back.colors[2-i][0] = h.c.down.colors[i][2]
		h.c.down.colors[i][2] = h.c.front.colors[i][2]
		h.c.front.colors[i][2] = tmp[i]
	}

	return h.c
}

type upRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Up() *upRotationHandler { return &upRotationHandler{h.c} }

func (h *upRotationHandler) Clockwise() *Rubikube {
	// Handle right rotation
	for i := 0; i < 2; i++ {
		h.c.up.colors[0][i], h.c.up.colors[2-i][0], h.c.up.colors[2][2-i], h.c.up.colors[i][2] =
			h.c.up.colors[2-i][0], h.c.up.colors[2][2-i], h.c.up.colors[i][2], h.c.up.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.front.colors[0][i]
		h.c.front.colors[0][i] = h.c.right.colors[0][i]
		h.c.right.colors[0][i] = h.c.back.colors[0][i]
		h.c.back.colors[0][i] = h.c.left.colors[0][i]
		h.c.left.colors[0][i] = tmp[i]
	}

	return h.c
}

func (h *upRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.up.colors[0][i], h.c.up.colors[i][2], h.c.up.colors[2][2-i], h.c.up.colors[2-i][0] =
			h.c.up.colors[i][2], h.c.up.colors[2][2-i], h.c.up.colors[2-i][0], h.c.up.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.front.colors[0][i]
		h.c.front.colors[0][i] = h.c.left.colors[0][i]
		h.c.left.colors[0][i] = h.c.back.colors[0][i]
		h.c.back.colors[0][i] = h.c.right.colors[0][i]
		h.c.right.colors[0][i] = tmp[i]
	}

	return h.c
}

type downRotationHandler struct{ c *Rubikube }

func (h *rotationHandler) Down() *downRotationHandler { return &downRotationHandler{h.c} }

func (h *downRotationHandler) Clockwise() *Rubikube {
	// Handle right rotation
	for i := 0; i < 2; i++ {
		h.c.down.colors[0][i], h.c.down.colors[2-i][0], h.c.down.colors[2][2-i], h.c.down.colors[i][2] =
			h.c.down.colors[2-i][0], h.c.down.colors[2][2-i], h.c.down.colors[i][2], h.c.down.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.front.colors[2][i]
		h.c.front.colors[2][i] = h.c.left.colors[2][i]
		h.c.left.colors[2][i] = h.c.back.colors[2][i]
		h.c.back.colors[2][i] = h.c.right.colors[2][i]
		h.c.right.colors[2][i] = tmp[i]
	}

	return h.c
}

func (h *downRotationHandler) CounterClockwise() *Rubikube {
	// Handle face rotation
	for i := 0; i < 2; i++ {
		h.c.down.colors[0][i], h.c.down.colors[i][2], h.c.down.colors[2][2-i], h.c.down.colors[2-i][0] =
			h.c.down.colors[i][2], h.c.down.colors[2][2-i], h.c.down.colors[2-i][0], h.c.down.colors[0][i]
	}

	// Handle sides rotation
	tmp := make([]color, 3)

	for i := 0; i < 3; i++ {
		tmp[i] = h.c.front.colors[2][i]
		h.c.front.colors[2][i] = h.c.right.colors[2][i]
		h.c.right.colors[2][i] = h.c.back.colors[2][i]
		h.c.back.colors[2][i] = h.c.left.colors[2][i]
		h.c.left.colors[2][i] = tmp[i]
	}

	return h.c
}
