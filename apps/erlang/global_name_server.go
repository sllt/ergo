package erlang

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
)

// TODO: https://github.com/erlang/otp/blob/master/lib/kernel/src/global.erl

type globalNameServer struct {
	gen.Server
}

// HandleCast
func (gns *globalNameServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}
