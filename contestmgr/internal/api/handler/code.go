package handler

import (
	"io"
	"net/http"

	"github.com/koss-null/toadstool/contestmgr/pkg/templater"
)

func Code(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodPut:
		{
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusAccepted)
			templater.NewTemplater()
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
