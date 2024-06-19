package main

import (
	"os"

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
	jsFileName := os.Args[1]

	wrappedVm, err := vm.NewVmWithFile(jsFileName)
	if err != nil {
		return err
	}
	wrappedVm.SetLogger(go_logger.Logger.CloneWithLevel("debug"))
	err = wrappedVm.RegisterModule(test_module.ModuleName, test_module.NewTestModuleModule(wrappedVm))
	if err != nil {
		return err
	}
	_, err = wrappedVm.Run()
	if err != nil {
		return err
	}

	return nil
}

// go run ./_example ./_example/js/main.js
