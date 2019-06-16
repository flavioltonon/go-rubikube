package rubikube

import (
	"math/rand"
)

type move interface {
	reverse() move
	instruction() instruction
}

type instruction string

const (
	INSTRUCTION_ROTATE_FRONT_CLOCKWISE        = "Rotate Front Clockwise"
	INSTRUCTION_ROTATE_FRONT_COUNTERCLOCKWISE = "Rotate Front CounterClockwise"
	INSTRUCTION_ROTATE_BACK_CLOCKWISE         = "Rotate Back Clockwise"
	INSTRUCTION_ROTATE_BACK_COUNTERCLOCKWISE  = "Rotate Back CounterClockwise"
	INSTRUCTION_ROTATE_LEFT_CLOCKWISE         = "Rotate Left Clockwise"
	INSTRUCTION_ROTATE_LEFT_COUNTERCLOCKWISE  = "Rotate Left CounterClockwise"
	INSTRUCTION_ROTATE_RIGHT_CLOCKWISE        = "Rotate Right Clockwise"
	INSTRUCTION_ROTATE_RIGHT_COUNTERCLOCKWISE = "Rotate Right CounterClockwise"
	INSTRUCTION_ROTATE_UP_CLOCKWISE           = "Rotate Up Clockwise"
	INSTRUCTION_ROTATE_UP_COUNTERCLOCKWISE    = "Rotate Up CounterClockwise"
	INSTRUCTION_ROTATE_DOWN_CLOCKWISE         = "Rotate Down Clockwise"
	INSTRUCTION_ROTATE_DOWN_COUNTERCLOCKWISE  = "Rotate Down CounterClockwise"
)

var standardMoves = map[instruction]move{
	INSTRUCTION_ROTATE_FRONT_CLOCKWISE:        frontRotateClockwise{},
	INSTRUCTION_ROTATE_FRONT_COUNTERCLOCKWISE: frontRotateCounterClockwise{},
	INSTRUCTION_ROTATE_BACK_CLOCKWISE:         backRotateClockwise{},
	INSTRUCTION_ROTATE_BACK_COUNTERCLOCKWISE:  backRotateCounterClockwise{},
	INSTRUCTION_ROTATE_LEFT_CLOCKWISE:         leftRotateClockwise{},
	INSTRUCTION_ROTATE_LEFT_COUNTERCLOCKWISE:  leftRotateCounterClockwise{},
	INSTRUCTION_ROTATE_RIGHT_CLOCKWISE:        rightRotateClockwise{},
	INSTRUCTION_ROTATE_RIGHT_COUNTERCLOCKWISE: rightRotateCounterClockwise{},
	INSTRUCTION_ROTATE_UP_CLOCKWISE:           upRotateClockwise{},
	INSTRUCTION_ROTATE_UP_COUNTERCLOCKWISE:    upRotateCounterClockwise{},
	INSTRUCTION_ROTATE_DOWN_CLOCKWISE:         downRotateClockwise{},
	INSTRUCTION_ROTATE_DOWN_COUNTERCLOCKWISE:  downRotateCounterClockwise{},
}

var moves = make([]move, 0)

func init() {
	for _, move := range standardMoves {
		moves = append(moves, move)
	}
}

type frontRotateClockwise struct{}

func (m frontRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_FRONT_CLOCKWISE }
func (m frontRotateClockwise) reverse() move            { return frontRotateCounterClockwise{} }

type frontRotateCounterClockwise struct{}

func (m frontRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_FRONT_COUNTERCLOCKWISE
}
func (m frontRotateCounterClockwise) reverse() move { return frontRotateClockwise{} }

type backRotateClockwise struct{}

func (m backRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_BACK_CLOCKWISE }
func (m backRotateClockwise) reverse() move            { return backRotateCounterClockwise{} }

type backRotateCounterClockwise struct{}

func (m backRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_BACK_COUNTERCLOCKWISE
}
func (m backRotateCounterClockwise) reverse() move { return backRotateClockwise{} }

type leftRotateClockwise struct{}

func (m leftRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_LEFT_CLOCKWISE }
func (m leftRotateClockwise) reverse() move            { return leftRotateCounterClockwise{} }

type leftRotateCounterClockwise struct{}

func (m leftRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_LEFT_COUNTERCLOCKWISE
}
func (m leftRotateCounterClockwise) reverse() move { return leftRotateClockwise{} }

type rightRotateClockwise struct{}

func (m rightRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_RIGHT_CLOCKWISE }
func (m rightRotateClockwise) reverse() move            { return rightRotateCounterClockwise{} }

type rightRotateCounterClockwise struct{}

func (m rightRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_RIGHT_COUNTERCLOCKWISE
}
func (m rightRotateCounterClockwise) reverse() move { return rightRotateClockwise{} }

type upRotateClockwise struct{}

func (m upRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_UP_CLOCKWISE }
func (m upRotateClockwise) reverse() move            { return upRotateCounterClockwise{} }

type upRotateCounterClockwise struct{}

func (m upRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_UP_COUNTERCLOCKWISE
}
func (m upRotateCounterClockwise) reverse() move { return upRotateClockwise{} }

type downRotateClockwise struct{}

func (m downRotateClockwise) instruction() instruction { return INSTRUCTION_ROTATE_DOWN_CLOCKWISE }
func (m downRotateClockwise) reverse() move            { return downRotateCounterClockwise{} }

type downRotateCounterClockwise struct{}

func (m downRotateCounterClockwise) instruction() instruction {
	return INSTRUCTION_ROTATE_DOWN_COUNTERCLOCKWISE
}
func (m downRotateCounterClockwise) reverse() move { return downRotateClockwise{} }

func newMove(instruction instruction) move { return standardMoves[instruction] }

func randomMove() move { return moves[rand.Intn(len(moves))] }
