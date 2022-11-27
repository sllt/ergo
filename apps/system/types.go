package system

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/lib"
)

type MessageSystemAnonMetrics struct {
	Name        string
	Arch        string
	OS          string
	NumCPU      int
	GoVersion   string
	ErgoVersion string
}

func RegisterTypes() error {
	types := []interface{}{
		MessageSystemAnonMetrics{},
	}
	rtOpts := etf.RegisterTypeOptions{Strict: true}

	for _, t := range types {
		if _, err := etf.RegisterType(t, rtOpts); err != nil && err != lib.ErrTaken {
			return err
		}
	}
	return nil
}
