package module

import (
	"github.com/dop251/goja"
	go_logger "github.com/pefish/go-logger"
)


type Console struct {
	Log func(data ...interface{}) `json:"log"`
}

var consoleModule = &Console{
	Log: func(data ...interface{}) {
		go_logger.Logger.Info(data...)
	},
}

type IWrappedVm interface {
	RegisterModule(moduleName string, module interface{}) error
	ToValue(i interface{}) goja.Value
}

func RegisterConsoleModule(v IWrappedVm) error {
	return v.RegisterModule("console", consoleModule)
}
