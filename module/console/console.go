package console

import (
	"github.com/pefish/go-jsvm/module"
)

const ModuleName = "console"
const ModuleName1 = "console_go"

type Console struct {
	vm module.IWrappedVm
}

func NewConsoleModule(vm module.IWrappedVm) *Console {
	return &Console{
		vm: vm,
	}
}

func (c *Console) Log(data ...interface{}) {
	c.Info(data...)
}

func (c *Console) Debug(data ...interface{}) {
	c.vm.Logger().Debug(data...)
}

func (c *Console) DebugF(format string, data ...interface{}) {
	c.vm.Logger().DebugF(format, data...)
}

func (c *Console) Info(data ...interface{}) {
	c.vm.Logger().Info(data...)
}

func (c *Console) InfoF(format string, data ...interface{}) {
	c.vm.Logger().InfoF(format, data...)
}

func (c *Console) Warn(data ...interface{}) {
	c.vm.Logger().Warn(data...)
}

func (c *Console) WarnF(format string, data ...interface{}) {
	c.vm.Logger().WarnF(format, data...)
}

func (c *Console) Error(data ...interface{}) {
	c.vm.Logger().Error(data...)
}

func (c *Console) ErrorF(format string, data ...interface{}) {
	c.vm.Logger().ErrorF(format, data...)
}
