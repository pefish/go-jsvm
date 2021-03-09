package command

import (
	"flag"
	"github.com/pefish/go-commander"
	vm2 "github.com/pefish/go-jsvm/pkg/vm"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
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
	fileInfo, err := os.Stat(jsFilename)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		return errors.New("illegal js file")
	}
	f, err := os.Open(jsFilename)
	if err != nil {
		return err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	vm, err := vm2.NewVmAndLoad(string(content))
	if err != nil {
		return err
	}
	_, err = vm.Run(nil)
	if err != nil {
		return err
	}
	return nil
}

