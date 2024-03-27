package main

import (
	"time"

	_ "github.com/dop251/goja_nodejs/console"
	vm "github.com/pefish/go-jsvm"
	"github.com/pefish/go-jsvm/_example/module/test_module"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	err := do()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}

func do() error {
	for {
		wrappedVm, err := vm.NewVmWithFile("./main.js")
		if err != nil {
			return err
		}
		err = wrappedVm.RegisterModule(test_module.ModuleName, test_module.NewTestModuleModule(wrappedVm))
		if err != nil {
			return err
		}
		_, err = wrappedVm.Run()
		if err != nil {
			return err
		}
		_, err = wrappedVm.RunMain(nil)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
