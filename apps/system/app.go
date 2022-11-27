package system

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/ergo/lib"
	"github.com/sllt/ergo/node"
)

func CreateApp(options node.System) gen.ApplicationBehavior {
	return &systemApp{
		options: options,
	}
}

type systemApp struct {
	gen.Application
	options node.System
}

func (sa *systemApp) Load(args ...etf.Term) (gen.ApplicationSpec, error) {
	lib.Log("SYSTEM: Application load")
	return gen.ApplicationSpec{
		Name:        "system_app",
		Description: "System Application",
		Version:     "v.1.0",
		Children: []gen.ApplicationChildSpec{
			gen.ApplicationChildSpec{
				Child: &systemAppSup{},
				Name:  "system_app_sup",
				Args:  []etf.Term{sa.options},
			},
		},
	}, nil
}

func (sa *systemApp) Start(p gen.Process, args ...etf.Term) {
	lib.Log("[%s] SYSTEM: Application started", p.NodeName())
}
