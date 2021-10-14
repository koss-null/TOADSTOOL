package handler

import (
	"io"
	"net/http"

	templater "github.com/koss-null/toadstool/contestmgr/pkg/tempalter"
)

const (
	TemplatePath = "../../../resources/compiler/solution_checker.mockgo"
	ResultPath   = "/usr/tmp"
)

func Code(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodPut:
		{
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusAccepted)

			// TODO: make a separate function
			go func() {
				t := templater.NewTemplater(TemplatePath, ResultPath)
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
