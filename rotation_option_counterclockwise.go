package rubikube

type rotationOptionCounterClockwise struct {
	cube *cube
}

func (r *rotationOptionCounterClockwise) Back() {
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
		tmp[i] = up[0][2-i]
		up[2][i] = right[i][0]
		right[i][0] = down[0][2-i]
		down[0][2-i] = left[2-i][2]
		left[2-i][2] = tmp[i]
	}
}

func (r *rotationOptionCounterClockwise) Front() {
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
