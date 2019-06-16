package rubikube

import "time"

type solveOption struct {
	*cube
}

func (s *solveOption) Optimized() (changes, time.Duration, error) {
	var start = time.Now()

	var solutionRoutine = make(changes, 0)

	var index int

	copy := s.cube.Copy()

	for copy.snapshot() != defaultSnapshot {
		var elapsed = time.Now().Sub(start)
		if elapsed.Seconds() > 10 {
			return changes{}, elapsed, ErrSolutionTookTooLong
		}

		var m move

		if data, exists := Solutions.Get(copy.snapshot().String()); exists {
			m = standardMoves[data.(*solution).NextStep.Instruction]
		} else {
			m = randomMove()
		}

		copy.apply(m)

		solutionRoutine = append(solutionRoutine, change{
			Index:       index + 1,
			Instruction: m.instruction(),
		})

		index++
	}

	for _, change := range solutionRoutine {
		s.cube.apply(standardMoves[change.Instruction])
	}

	return solutionRoutine, time.Now().Sub(start), nil
}

func (s *solveOption) Randomically() (changes, time.Duration, error) {
	var start = time.Now()

	var solutionRoutine = make(changes, 0)

	var index int

	copy := s.cube.Copy()

	for copy.snapshot() != defaultSnapshot {
		var elapsed = time.Now().Sub(start)
		if elapsed.Seconds() > 30 {
			return changes{}, elapsed, ErrSolutionTookTooLong
		}

		var m = randomMove()

		copy.apply(m)

		solutionRoutine = append(solutionRoutine, change{
			Index:       index + 1,
			Instruction: m.instruction(),
		})

		index++
	}

	for _, change := range solutionRoutine {
		s.cube.apply(standardMoves[change.Instruction])
	}

	return solutionRoutine, time.Now().Sub(start), nil
}
