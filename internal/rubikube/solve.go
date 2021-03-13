package rubikube

import (
	"fmt"
)

var solutions map[string]bool

func init() {
	solutions = make(map[string]bool)

	for _, s := range buildSolutions() {
		solutions[s] = true
	}
}

type Solution struct {
	movements []movement
}

func (s *Solution) Print() {
	for i, movement := range s.movements {
		fmt.Printf("%d: %s\n", i+1, movement.String())
	}
}

func (c *Rubikube) Solve() *Solution {
	q := newQueue()

	q.enqueue(&node{
		cube: c.copy(),
		solution: &Solution{
			movements: make([]movement, 0),
		},
	})

	visited := make(map[string]bool)

	for q.len() > 0 {
		currentNode := q.dequeue()

		ss := currentNode.cube.snapshot()

		if _, isSolution := solutions[ss]; isSolution {
			return currentNode.solution
		}

		if _, isVisited := visited[ss]; isVisited {
			continue
		}

		visited[ss] = true

		nodes := make(chan *node, len(movements))

		go func() {
			for i := range movements {
				next := currentNode.cube.copy()

				next.do(movements[i])

				nodes <- &node{
					cube: next,
					solution: &Solution{
						movements: append(currentNode.solution.movements, movements[i]),
					},
				}
			}

			close(nodes)
		}()

		for i := 0; i < len(movements); i++ {
			select {
			case node := <-nodes:
				q.enqueue(node)
			}
		}
	}

	return nil
}

type node struct {
	cube     *Rubikube
	solution *Solution

	next     *node
	previous *node
}

type queue struct {
	head   *node
	tail   *node
	length int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) enqueue(n *node) {
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		n.next = q.head
		q.head.previous = n
		q.head = n
	}

	q.length++
}

func (q *queue) dequeue() *node {
	n := q.tail

	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail.previous.next = nil
		q.tail = q.tail.previous
	}

	q.length--

	return n
}

func (q *queue) len() int {
	return q.length
}

func buildSolutions() []string {
	return combine([]string{
		"111111111",
		"222222222",
		"333333333",
		"444444444",
		"555555555",
		"666666666",
	})
}

func combine(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	}

	if len(strings) == 1 {
		return strings
	}

	combinations := make([]string, 0, factorial(len(strings)))

	for i := range strings {
		strings[0], strings[i] = strings[i], strings[0]

		for _, s := range combine(strings[1:]) {
			combinations = append(combinations, fmt.Sprintf("%s%s", strings[0], s))
		}
	}

	return combinations
}

func factorial(n int) int {
	if n < 1 {
		panic("n should be higher than zero")
	}

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
