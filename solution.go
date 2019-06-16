package rubikube

import (
	"fmt"
	"os"
	"sync"

	mcache "github.com/OrlovEvgeny/go-mcache"
)

type solution struct {
	Snapshot       snapshot `bson:"snapshot"`
	NextStep       *step    `bson:"nextStep"`
	RemainingSteps *int     `bson:"remainingSteps"`
}

type step struct {
	Snapshot    snapshot    `bson:"snapshot"`
	Instruction instruction `bson:"instruction"`
}

//Start mcache instance
var Solutions = mcache.New()

type option int

const (
	EXIT option = iota
	SHUFFLE
	PRINT
	SOLVE
	COUNT_SOLUTIONS
	RESET
)

type printSuboption int

const (
	PRINT_RETURN printSuboption = iota
	PRINT_TRIDIMENSIONAL
	PRINT_BIDIMENSIONAL
)

type printBidimensionalSuboption int

const (
	PRINT_BIDIMENSIONAL_RETURN printBidimensionalSuboption = iota
	PRINT_BIDIMENSIONAL_FRONT_FACE
	PRINT_BIDIMENSIONAL_BACK_FACE
	PRINT_BIDIMENSIONAL_LEFT_FACE
	PRINT_BIDIMENSIONAL_RIGHT_FACE
	PRINT_BIDIMENSIONAL_UP_FACE
	PRINT_BIDIMENSIONAL_DOWN_FACE
)

type solveSuboption int

const (
	SOLVE_RETURN solveSuboption = iota
	SOLVE_OPTIMIZED
	SOLVE_RANDOMICALLY
)

func Run() error {
	go buildSolutions()

	var cube = NewCube()

	for {
		var option option

		fmt.Print("---------- MAIN MENU ----------\n\n")
		fmt.Print("Choose one of the options below:\n\n1) Shuffle cube\n2) Print cube\n3) Solve cube\n4) Count current solutions\n5) Reset cube\n0) Exit\n\nOption: ")
		_, err := fmt.Scan(&option)
		if err != nil {
			return err
		}

		fmt.Print("\n")

		switch option {
		case SHUFFLE:
			err = cube.handleShuffleOption()
			if err != nil {
				return err
			}
		case PRINT:
			err = cube.handlePrintOption()
			if err != nil {
				return err
			}
		case SOLVE:
			err = cube.handleSolveOption()
			if err != nil {
				return err
			}
		case COUNT_SOLUTIONS:
			err = cube.handleCountSolutionsOption()
			if err != nil {
				return err
			}
		case RESET:
			cube = NewCube()
		case EXIT:
			os.Exit(0)
		default:
			fmt.Print("Invalid option!\n")
		}

		fmt.Print("\n")
	}
}

func buildSolutions() error {
	for size := 0; size < 15; size++ {

		err := findSolutions(size)
		if err == ErrBulkLowerSolutionNotFound {
			size -= 2
			continue
		}
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func findSolutions(size int) error {
	if size == 0 {
		if _, exists := Solutions.Get(defaultSnapshot.String()); exists {
			return ErrSolutionsAlreadyExists
		}

		return Solutions.Set(
			defaultSnapshot.String(),
			&solution{
				Snapshot:       defaultSnapshot,
				RemainingSteps: &size,
			},
			mcache.TTL_FOREVER,
		)
	}

	var stepsUpperLimit = GODS_NUMBER
	if size < stepsUpperLimit {
		stepsUpperLimit = size
	}

	var alreadyMappedRecurrency int
	var lowerSolutionNotFound int
	var queue = make(chan bool, 1)
	var rwm sync.RWMutex

	for {
		queue <- true

		if alreadyMappedRecurrency >= 20 {
			return nil
		}
		if lowerSolutionNotFound >= 50 {
			return ErrBulkLowerSolutionNotFound
		}

		go func() {
			defer func() {
				<-queue
			}()

			var cube = NewCube()

			cube.Shuffle(size)

			// Solution has already been mapped
			if _, exists := Solutions.Get(cube.snapshot().String()); exists {
				alreadyMappedRecurrency++
				return
			}

			var innerWg sync.WaitGroup
			var newSolutionFound bool

			for _, m := range standardMoves {
				innerWg.Add(1)

				go func(m move) {
					defer innerWg.Done()

					copy := cube.Copy()

					copy.apply(m)

					if _, exists := Solutions.Get(copy.snapshot().String()); !exists {
						return
					}

					newSolutionFound = true

					// Handle new solution
					rwm.Lock()

					Solutions.Set(
						cube.snapshot().String(),
						&solution{
							Snapshot: cube.snapshot(),
							NextStep: &step{
								Snapshot:    copy.snapshot(),
								Instruction: m.instruction(),
							},
							RemainingSteps: &stepsUpperLimit,
						},
						mcache.TTL_FOREVER,
					)

					rwm.Unlock()
				}(m)
			}

			innerWg.Wait()

			if newSolutionFound {
				alreadyMappedRecurrency = 0
				lowerSolutionNotFound = 0
				return
			}

			lowerSolutionNotFound++
		}()
	}
}

func (c *cube) handleShuffleOption() error {
	var n int

	fmt.Print("Input the number of times the cube\nshould be shuffled\n\nTimes: ")
	fmt.Scan(&n)

	c.Shuffle(n)

	return nil
}

func (c cube) handlePrintOption() error {
	var s printSuboption

	for {
		fmt.Print("--------- PRINT  MENU ---------\n\n")
		fmt.Print("Choose one of the print types below:\n1) Tridimensional\n2) Bidimensional\n0) Return\n\nOption: ")
		fmt.Scan(&s)

		switch s {
		case PRINT_TRIDIMENSIONAL:
			c.Front().Print().Tridimensional()
		case PRINT_BIDIMENSIONAL:
			c.handlePrintBidimensionalSuboption()
		case PRINT_RETURN:
			return nil
		default:
			fmt.Print("Invalid option!\n")
		}

		fmt.Print("\n")
	}
}

func (c cube) handlePrintBidimensionalSuboption() {
	var f printBidimensionalSuboption

	for {
		fmt.Print("Choose one of the faces below:\n1) Front\n2) Back\n3) Left\n4) Right\n5) Up\n6) Down\n0) Return\n\nOption: ")
		fmt.Scan(&f)

		switch f {
		case PRINT_BIDIMENSIONAL_FRONT_FACE:
			c.Front().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_BACK_FACE:
			c.Back().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_LEFT_FACE:
			c.Left().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_RIGHT_FACE:
			c.Right().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_UP_FACE:
			c.Up().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_DOWN_FACE:
			c.Down().Print().Bidimensional()
		case PRINT_BIDIMENSIONAL_RETURN:
			return
		default:
			fmt.Print("Invalid option!\n")
		}

		fmt.Print("\n")
	}
}

func (c cube) handleSolveOption() error {
	var s solveSuboption

	for {
		fmt.Print("------- SOLUTIONS  MENU -------\n\n")
		fmt.Print("Choose one of the solution methods below:\n1) Optimized\n2) Randomically\n0) Return\n\nOption: ")
		fmt.Scan(&s)
		fmt.Print("\n")

		switch s {
		case SOLVE_OPTIMIZED:
			solution, duration, err := c.Solve().Optimized()
			if err == ErrSolutionTookTooLong {
				fmt.Print("The solution is taking more time than expected to be found. Try again after a while.")
				return nil
			}
			if err != nil {
				return err
			}

			if duration.Seconds() < 1 {
				fmt.Print(fmt.Sprintf("The cube input took %v nanoseconds to be solved.\n\n", duration.Nanoseconds()))
			} else {
				fmt.Print(fmt.Sprintf("The cube input took %v seconds to be solved.\n\n", duration.Seconds()))
			}

			fmt.Print("Steps to solution:\n")
			for _, step := range solution {
				fmt.Print(fmt.Sprintf("%v: %v\n", step.Index, step.Instruction))
			}

			return nil
		case SOLVE_RANDOMICALLY:
			solution, duration, err := c.Solve().Randomically()
			if err == ErrSolutionTookTooLong {
				fmt.Print("The solution is taking more time than expected to be found. Try again after a while.\n")
				return nil
			}
			if err != nil {
				return err
			}

			if duration.Seconds() < 1 {
				fmt.Print(fmt.Sprintf("The cube input took %v nanoseconds to be solved.\n\n", duration.Nanoseconds()))
			} else {
				fmt.Print(fmt.Sprintf("The cube input took %v seconds to be solved.\n\n", duration.Seconds()))
			}

			fmt.Print("Steps to solution:\n")
			for _, step := range solution {
				fmt.Print(fmt.Sprintf("%v: %v\n", step.Index, step.Instruction))
			}

			return nil
		case SOLVE_RETURN:
			return nil
		default:
			fmt.Print("Invalid option!\n")
		}

		fmt.Print("\n")
	}
}

func (c cube) handleCountSolutionsOption() error {
	fmt.Print(fmt.Sprintf("%d solutions have been found so far.\n\n", Solutions.Len()))

	return nil
}
