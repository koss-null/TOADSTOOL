package api

import (
	"net/http"

	"github.com/koss-null/toadstool/contestmgr/internal/api/handler"
)

const (
	CodePath = "/code"
)

func StartCodeServer() error {
	// TODO: we only need to handle the code if we are not checking any other code
	http.HandleFunc(CodePath, handler.Code)
	return nil
}
