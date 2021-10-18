package checker

import (
	"github.com/koss-null/toadstool/contestmgr/pkg/solution"
	. "github.com/koss-null/toadstool/contestmgr/pkg/utils/queue"
)

type Checker struct {
	solutions Queue
	results   Queue

	stop chan interface{}
}

func (c *Checker) AddSolution(s *solution.Solution) {
	c.solutions.Push(*s)
}

func (c *Checker) checkNextSolution() {
	sol := c.solutions.Pop().(*solution.Solution)
	// TODO: check the solution
	c.results.Push(result)
}

func (c *Checker) Run() {
	c.solutions = solution.NewQueue()
	c.results = result.NewQueue()
	c.stop = make(chan interface{})
	for {
		select {
		case <-c.stop:
			return
		default:
			c.checkNextSolution()
		}
	}
}

func (c Checker) Stop() {
	close(c.stop)
}
