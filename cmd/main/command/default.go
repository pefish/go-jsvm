package command

import (
	"flag"
	"github.com/pefish/go-commander"
	vm2 "github.com/pefish/go-jsvm/pkg/vm"
	"github.com/pkg/errors"
)

type DefaultCommand struct {

}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{

	}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	if len(data.Args) == 0 {
		return errors.New("please set js file")
	}
	jsFilename := data.Args[0]
	vm, err := vm2.NewVmAndLoadWithFile(jsFilename)
	if err != nil {
		return err
	}
	_, err = vm.Run(nil)
	if err != nil {
		return err
	}
	return nil
}

