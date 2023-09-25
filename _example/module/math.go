package module

import (
	"github.com/pefish/go-jsvm/pkg/vm/module"
	"math"
)

type Math struct {
	vm module.IWrappedVm
}

func (c *Math) Abs(a float64) float64 {
	return math.Abs(a)
}

func NewMathModule(vm module.IWrappedVm) *Math {
	return &Math{
		vm: vm,
	}
}
