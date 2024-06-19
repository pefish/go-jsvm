package time

import (
	"time"

	"github.com/pefish/go-jsvm/module"
)

const ModuleName = "time_go"

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

func (c *Time) SetInterval(seconds int, func_ func() bool) {
	timer := time.NewTimer(0)
	for range timer.C {
		exit := func_()
		if exit {
			return
		}
		timer.Reset(time.Duration(seconds) * time.Second)
	}
}
