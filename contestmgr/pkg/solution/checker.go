package solution

import (
	"github.com/koss-null/toadstool/contestmgr/internal/codewriter"
	templater "github.com/koss-null/toadstool/contestmgr/pkg/templater"
	. "github.com/koss-null/toadstool/contestmgr/pkg/utils/queue"
)

const (
	TEMPLATE_PATH = "../../../resources/compiler/solution_checker.mockgo"
	RESULT_PATH   = "/usr/tmp"

	ERROR_BUFFER_SIZE = 100
)

type Checker struct {
	solutions Queue
	results   Queue

	writer codewriter.Writer
	stop   chan error
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

	var mainCode string
	var err error
	if mainCode, err = t.Build(); err != nil {
		// FIXME: there is a lot of functions that can return a error, need to handle
		// the error pushing mechanism
		return
	}
	c.writer.Write(codewriter.MAIN_FILE, mainCode)
	// TODO: sol.Code need to be prepared for importing first
	c.writer.Write(codewriter.SOLUTION_FILE, sol.Code)
	// TODO: change TaskID into the TaskID tests
	c.writer.WriteBuffered(codewriter.TEST_FILE, sol.TaskID)

	// TODO: run code

	c.results.Push(result)
}

// returns a chan which need to be stopped to stop checker running
// any errors that occurs during the Run() are stored in thin chan
func (c *Checker) Run() chan error {
	c.solutions = NewQueue()
	c.results = result.NewQueue()
	c.stop = make(chan error, ERROR_BUFFER_SIZE)

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
