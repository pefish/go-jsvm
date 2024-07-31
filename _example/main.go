package main

import (
	"log"
	"os"
	"time"

	_ "github.com/dop251/goja_nodejs/console"
	vm "github.com/pefish/go-jsvm"
	"github.com/pefish/go-jsvm/_example/module/test_module"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	jsFileName := os.Args[1]

	wrappedVm, err := vm.NewVmWithFile(jsFileName)
	if err != nil {
		return err
	}
	err = wrappedVm.RegisterModule(test_module.ModuleName, test_module.NewTestModuleModule(wrappedVm))
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(10 * time.Second)
		wrappedVm.Kill()
	}()

	_, err = wrappedVm.Run()
	if err != nil {
		return err
	}

	return nil
}

// go run ./_example ./_example/js/main.js
