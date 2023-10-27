package main

import (
	_ "github.com/dop251/goja_nodejs/console"
	vm "github.com/pefish/go-jsvm"
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
	wrappedVm, err := vm.NewVmWithFile("./main.js")
	if err != nil {
		return err
	}

	for {
		_, err := wrappedVm.Run()
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
