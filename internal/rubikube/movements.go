package rubikube

type movement int

const (
	ROTATE_FRONT_CLOCKWISE movement = iota + 1
	ROTATE_FRONT_COUNTERCLOCKWISE
	ROTATE_BACK_CLOCKWISE
	ROTATE_BACK_COUNTERCLOCKWISE
	ROTATE_LEFT_CLOCKWISE
	ROTATE_LEFT_COUNTERCLOCKWISE
	ROTATE_RIGHT_CLOCKWISE
	ROTATE_RIGHT_COUNTERCLOCKWISE
	ROTATE_UP_CLOCKWISE
	ROTATE_UP_COUNTERCLOCKWISE
	ROTATE_DOWN_CLOCKWISE
	ROTATE_DOWN_COUNTERCLOCKWISE
)

var movements = []movement{
	ROTATE_FRONT_CLOCKWISE,
	ROTATE_FRONT_COUNTERCLOCKWISE,
	ROTATE_BACK_CLOCKWISE,
	ROTATE_BACK_COUNTERCLOCKWISE,
	ROTATE_LEFT_CLOCKWISE,
	ROTATE_LEFT_COUNTERCLOCKWISE,
	ROTATE_RIGHT_CLOCKWISE,
	ROTATE_RIGHT_COUNTERCLOCKWISE,
	ROTATE_UP_CLOCKWISE,
	ROTATE_UP_COUNTERCLOCKWISE,
	ROTATE_DOWN_CLOCKWISE,
	ROTATE_DOWN_COUNTERCLOCKWISE,
}

func (m movement) String() string {
	switch m {
	case ROTATE_FRONT_CLOCKWISE:
		return "Rotate Front Clockwise"
	case ROTATE_FRONT_COUNTERCLOCKWISE:
		return "Rotate Front CounterClockwise"
	case ROTATE_BACK_CLOCKWISE:
		return "Rotate Back Clockwise"
	case ROTATE_BACK_COUNTERCLOCKWISE:
		return "Rotate Back CounterClockwise"
	case ROTATE_LEFT_CLOCKWISE:
		return "Rotate Left Clockwise"
	case ROTATE_LEFT_COUNTERCLOCKWISE:
		return "Rotate Left CounterClockwise"
	case ROTATE_RIGHT_CLOCKWISE:
		return "Rotate Right Clockwise"
	case ROTATE_RIGHT_COUNTERCLOCKWISE:
		return "Rotate Right CounterClockwise"
	case ROTATE_UP_CLOCKWISE:
		return "Rotate Up Clockwise"
	case ROTATE_UP_COUNTERCLOCKWISE:
		return "Rotate Up CounterClockwise"
	case ROTATE_DOWN_CLOCKWISE:
		return "Rotate Down Clockwise"
	case ROTATE_DOWN_COUNTERCLOCKWISE:
		return "Rotate Down CounterClockwise"
	default:
		return ""
	}
}
