package main

import (
	"github.com/sllt/ergo/gen"
)

type rootHandler struct {
	gen.WebHandler
}

func (r *rootHandler) HandleRequest(process *gen.WebHandlerProcess, request gen.WebMessageRequest) gen.WebHandlerStatus {
	request.Response.Write([]byte("Hello"))
	return gen.WebHandlerStatusDone
}
