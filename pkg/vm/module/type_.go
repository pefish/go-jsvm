package module

import "github.com/dop251/goja"

type IWrappedVm interface {
	RegisterModule(moduleName string, module interface{}) error
	ToValue(i interface{}) goja.Value
}
