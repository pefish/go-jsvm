package main

import (
	"github.com/pefish/go-jsvm/pkg/vm"
	"go-jsvm/_example/module"
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
	wrappedVm, err := vm.NewVm(`
function main() {
  console.log(math.abs(-1))
}
`)
	if err != nil {
		return err
	}
	err = wrappedVm.RegisterModule("math", module.NewMathModule(wrappedVm))
	if err != nil {
		return err
	}

	for {
		_, err := wrappedVm.Run(nil)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
