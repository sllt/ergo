package system

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
)

type systemAppSup struct {
	gen.Supervisor
}

func (sas *systemAppSup) Init(args ...etf.Term) (gen.SupervisorSpec, error) {
	return gen.SupervisorSpec{
		Children: []gen.SupervisorChildSpec{
			gen.SupervisorChildSpec{
				Name:  "system_metrics",
				Child: &systemMetrics{},
				Args:  args,
			},
		},
		Strategy: gen.SupervisorStrategy{
			Type:      gen.SupervisorStrategyOneForOne,
			Intensity: 10,
			Period:    5,
			Restart:   gen.SupervisorStrategyRestartPermanent,
		},
	}, nil
}
