package console

import (
	"github.com/pefish/go-jsvm/module"
)

const ModuleName = "console"
const ModuleName1 = "console_go"

type Console struct {
	vm       module.IWrappedVm
	hookFunc func(data ...interface{})
}

func NewConsoleModule(vm module.IWrappedVm) *Console {
	return &Console{
		vm: vm,
	}
}

func (c *Console) SetHookFunc(func_ func(data ...interface{})) {
	c.hookFunc = func_
}

func (c *Console) Log(data ...interface{}) {
	if c.hookFunc != nil {
		c.hookFunc(data...)
		return
	}
	c.vm.Logger().Info(data...)
}
