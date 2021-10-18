package handler

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/koss-null/toadstool/contestmgr/pkg/solution"
)

const (
	TASK_ID_HEADER = "X-Solution-Id"
)

var once sync.Once
var checker solution.Checker

func Init() {
	once.Do(func() {
		checker = solution.Checker{}
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

			checker.AddSolution(&solution.Solution{
				Code:   string(body),
				TaskID: strings.Join(taskID, ""),
			})
			w.WriteHeader(http.StatusAccepted)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
