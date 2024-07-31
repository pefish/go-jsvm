package module

import (
	"github.com/dop251/goja"
	i_logger "github.com/pefish/go-interface/i-logger"
)

type IWrappedVm interface {
	RegisterModule(moduleName string, module interface{}) error
	ToValue(i interface{}) goja.Value
	Logger() i_logger.ILogger
	Panic(err error)
	PanicWithMsg(msg string)
}
