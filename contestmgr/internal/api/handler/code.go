package handler

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/koss-null/toadstool/contestmgr/pkg/solution"
	templater "github.com/koss-null/toadstool/contestmgr/pkg/tempalter"
	. "github.com/koss-null/toadstool/contestmgr/pkg/utils/queue"
)

const (
	// FIXME: put it in Docker and use an absolute path
	TEMPLATE_PATH = "../../../resources/compiler/solution_checker.mockgo"
	RESULT_PATH   = "/usr/tmp"

	TASK_ID_HEADER = "Solution-ID"
)

var once sync.Once
var queue Queue

func Init() {
	once.Do(func() {
		queue = solution.NewQueue()
	})
}

func Code(w http.ResponseWriter, r *http.Request) {
	Init()
	defer r.Body.Close()
	switch r.Method {
	case http.MethodPut:
		{
			// searching for "Task-ID" header
			var taskID []string
			var ok bool
			if taskID, ok = r.Header[TASK_ID_HEADER]; !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// reading the request body (should contain only user's code)
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			queue.Push(solution.Solution{
				Code:   string(body),
				TaskID: strings.Join(taskID, ""),
			})
			w.WriteHeader(http.StatusAccepted)

			// TODO: make a separate function
			go func() {
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
			}()
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
