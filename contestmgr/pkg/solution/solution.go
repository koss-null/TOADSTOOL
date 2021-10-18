package solution

import (
	"container/list"
	"fmt"
	"sync"

	. "github.com/koss-null/toadstool/contestmgr/pkg/utils/queue"
)

type Solution struct {
	Code   string
	TaskID string
}

type (
	solutionQueue struct {
		mutex     *sync.Mutex
		solutions *list.List
	}
)

func NewQueue() Queue {
	return &solutionQueue{
		mutex:     &sync.Mutex{},
		solutions: list.New(),
	}
}

func (q *solutionQueue) Push(s interface{}) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	switch s.(type) {
	case Solution:
		q.solutions.PushBack(s)
	default:
		return fmt.Errorf(
			"forribden trying to add type '%T' into solution list",
			s,
		)
	}

	return nil
}

func (q *solutionQueue) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	el := q.solutions.Front()
	if el == nil {
		return nil
	}
	q.solutions.Remove(el)
	return el.Value
}
