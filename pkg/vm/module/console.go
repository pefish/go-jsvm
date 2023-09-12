package module

type Console struct {
	vm IWrappedVm
}

func (c *Console) Log(data ...interface{}) {
	c.vm.Logger().Info(data...)
}

func NewConsoleModule(vm IWrappedVm) *Console {
	return &Console{
		vm: vm,
	}
}
