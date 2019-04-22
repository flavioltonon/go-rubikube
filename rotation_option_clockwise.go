package rubikube

type rotationOptionClockwise struct {
	cube *cube
}

func (r *rotationOptionClockwise) Front() {
	var (
		front = r.cube.front.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)

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
}

func (r *rotationOptionClockwise) Back() {
	var (
		back  = r.cube.back.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)

	// Handle back rotation
	for i := 0; i < 2; i++ {
		back[0][i], back[i][2], back[2][2-i], back[2-i][0] = back[i][2-i], back[2][2-i], back[2-i][0], back[0][i]
	}

	// Handle sides rotation
	for i := 0; i < 3; i++ {
		tmp[i] = up[0][i]
		up[0][i] = right[i][2]
		right[i][2] = down[2][2-i]
		down[2][2-i] = left[i][0]
		left[i][0] = tmp[i]
	}
}

func (r *rotationOptionClockwise) Left() {
	var (
		front = r.cube.front.colors
		back  = r.cube.back.colors
		left  = r.cube.left.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)
	// Handle left rotation
	for i := 0; i < 2; i++ {
		left[0][i], left[i][2], left[2][2-i], left[2-i][0] = left[i][2-i], left[2][2-i], left[2-i][0], left[0][i]
	}

	// Handle sides rotation
	for i := 0; i < 3; i++ {
		tmp[i] = up[i][0]
		up[i][0] = back[2-i][2]
		back[2-i][2] = down[i][0]
		down[i][0] = front[i][0]
		front[i][0] = tmp[i]
	}
}

func (r *rotationOptionClockwise) Right() {
	var (
		front = r.cube.front.colors
		back  = r.cube.back.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)
	// Handle right rotation
	for i := 0; i < 2; i++ {
		right[0][i], right[i][2], right[2][2-i], right[2-i][0] = right[i][2-i], right[2][2-i], right[2-i][0], right[0][i]
	}

	// Handle sides rotation
	for i := 0; i < 3; i++ {
		tmp[i] = up[i][2]
		up[i][2] = front[i][2]
		front[i][2] = down[i][2]
		down[i][2] = back[2-i][0]
		back[2-i][0] = tmp[i]
	}
}

func (r *rotationOptionClockwise) Up() {
	var (
		front = r.cube.front.colors
		back  = r.cube.back.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		up    = r.cube.up.colors
		tmp   = make([]color, 3)
	)
	// Handle right rotation
	for i := 0; i < 2; i++ {
		up[0][i], up[i][2], up[2][2-i], up[2-i][0] = up[i][2-i], up[2][2-i], up[2-i][0], up[0][i]
	}

	// Handle sides rotation
	for i := 0; i < 3; i++ {
		tmp[i] = front[0][i]
		front[0][i] = right[0][i]
		right[0][i] = back[0][i]
		back[0][i] = left[0][i]
		left[0][i] = tmp[i]
	}
}

func (r *rotationOptionClockwise) Down() {
	var (
		front = r.cube.front.colors
		back  = r.cube.back.colors
		left  = r.cube.left.colors
		right = r.cube.right.colors
		down  = r.cube.down.colors
		tmp   = make([]color, 3)
	)
	// Handle right rotation
	for i := 0; i < 2; i++ {
		down[0][i], down[i][2], down[2][2-i], down[2-i][0] = down[i][2-i], down[2][2-i], down[2-i][0], down[0][i]
	}

	// Handle sides rotation
	for i := 0; i < 3; i++ {
		tmp[i] = front[2][i]
		front[2][i] = left[2][i]
		left[2][i] = back[2][i]
		back[2][i] = right[2][i]
		right[2][i] = tmp[i]
	}
}
