package solution

import (
	templater "github.com/koss-null/toadstool/contestmgr/pkg/templater"
	. "github.com/koss-null/toadstool/contestmgr/pkg/utils/queue"
)

const (
	TEMPLATE_PATH = "../../../resources/compiler/solution_checker.mockgo"
	RESULT_PATH   = "/usr/tmp"
)

type Checker struct {
	solutions Queue
	results   Queue

	stop chan interface{}
}

func (c *Checker) AddSolution(s *Solution) {
	c.solutions.Push(*s)
}

func (c *Checker) checkNextSolution() {
	sol := c.solutions.Pop().(*Solution)
	// TODO: check the solution

	t := templater.NewTemplater(TEMPLATE_PATH, RESULT_PATH)
	// TODO: add all handlers
	t.AddHandler()
	// TODO: add the code file to the ResultPath (it's in body var)
	if err := t.Build(); err != nil {
		// FIXME: there is a lot of functions that can return a error, need to handle
		// the error pushing mechanism
		return
	}
	// TODO: compile code
	// TODO: run code

	c.results.Push(result)
}

// returns a chan which need to be stopped to stop checker running
func (c *Checker) Run() chan interface{} {
	c.solutions = NewQueue()
	c.results = result.NewQueue()
	c.stop = make(chan interface{})

	go func() {
		for {
			select {
			case <-c.stop:
				return
			default:
				c.checkNextSolution()
			}
		}
	}()

	return c.stop
}

func (c Checker) Stop() {
	close(c.stop)
}
