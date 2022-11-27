package cloud

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
)

type cloudAppSup struct {
	gen.Supervisor
}

func (cas *cloudAppSup) Init(args ...etf.Term) (gen.SupervisorSpec, error) {
	return gen.SupervisorSpec{
		Children: []gen.SupervisorChildSpec{
			gen.SupervisorChildSpec{
				Name:  "cloud_client",
				Child: &cloudClient{},
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
