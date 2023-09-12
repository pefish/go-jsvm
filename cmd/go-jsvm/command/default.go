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
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Init(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	jsFilename, ok := data.Args["js file"]
	if !ok {
		return errors.New("please set js file")
	}
	vm, err := vm2.NewVmWithFile(jsFilename)
	if err != nil {
		return err
	}
	_, err = vm.Run(nil)
	if err != nil {
		return err
	}
	return nil
}
