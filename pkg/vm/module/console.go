package module

import (
	go_logger "github.com/pefish/go-logger"
)

type Console struct {
	Log func(data ...interface{}) `json:"log"`
}

func RegisterConsoleModule(v IWrappedVm) error {
	return v.RegisterModule("console", &Console{
		Log: func(data ...interface{}) {
			go_logger.Logger.Info(data...)
		},
	})
}
