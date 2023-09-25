package command

import (
	"flag"
	"github.com/pefish/go-commander"
	vm "github.com/pefish/go-jsvm"
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
	wrappedVm, err := vm.NewVmWithFile(jsFilename)
	if err != nil {
		return err
	}
	_, err = wrappedVm.Run(nil)
	if err != nil {
		return err
	}
	return nil
}
