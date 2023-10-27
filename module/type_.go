package module

import (
	"github.com/dop251/goja"
	go_logger "github.com/pefish/go-logger"
)

type IWrappedVm interface {
	RegisterModule(moduleName string, module interface{}) error
	ToValue(i interface{}) goja.Value
	Logger() go_logger.InterfaceLogger
	Panic(err error)
}
