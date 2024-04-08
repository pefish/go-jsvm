package time

import (
	"time"

	"github.com/pefish/go-jsvm/module"
)

const ModuleName = "time"

type Time struct {
	vm module.IWrappedVm
}

func NewTimeModule(vm module.IWrappedVm) *Time {
	return &Time{
		vm: vm,
	}
}

func (c *Time) Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
