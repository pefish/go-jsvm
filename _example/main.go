package main

import (
	_ "github.com/dop251/goja_nodejs/console"
	vm "github.com/pefish/go-jsvm"
	"github.com/pefish/go-jsvm/_example/module/test_module"
	"log"
	"time"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
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
		time.Sleep(2 * time.Second)
	}
}
