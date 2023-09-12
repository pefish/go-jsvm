package vm

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/pefish/go-error"
	"github.com/pefish/go-jsvm/pkg/vm/module"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

type WrappedVm struct {
	Vm     *goja.Runtime
	script string
}

type MainFuncType func([]interface{}) interface{}

func NewVmAndLoad(script string) (*WrappedVm, error) {
	vm, err := NewVm(script)
	if err != nil {
		return nil, err
	}
	err = vm.Load()
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func NewVmAndLoadWithFile(jsFilename string) (*WrappedVm, error) {
	fileInfo, err := os.Stat(jsFilename)
	if err != nil {
		return nil, go_error.WithStack(err)
	}
	if fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		return nil, errors.New("illegal js file")
	}
	f, err := os.Open(jsFilename)
	if err != nil {
		return nil, go_error.WithStack(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, go_error.WithStack(err)
	}
	vm, err := NewVmAndLoad(string(content))
	if err != nil {
		return nil, go_error.WithStack(err)
	}
	return vm, nil
}

func NewVm(script string) (*WrappedVm, error) {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	wrappedVm := &WrappedVm{
		Vm:     vm,
		script: script,
	}
	err := wrappedVm.registerModules()
	if err != nil {
		return nil, err
	}
	return wrappedVm, nil
}

// 注册预设的一些模块
func (v *WrappedVm) registerModules() error {
	err := module.RegisterConsoleModule(v)
	if err != nil {
		return err
	}
	err = module.RegisterRegexModule(v)
	if err != nil {
		return err
	}

	return nil
}

func (v *WrappedVm) RegisterModule(moduleName string, module interface{}) error {
	err := v.Vm.Set(moduleName, module)
	if err != nil {
		return go_error.WithStack(err)
	}
	return nil
}

func (v *WrappedVm) ToValue(i interface{}) goja.Value {
	return v.Vm.ToValue(i)
}

// 执行脚本
func (v *WrappedVm) Load() error {
	_, err := v.Vm.RunString(v.script)
	if err != nil {
		return go_error.WithStack(err)
	}
	return nil
}

// 执行脚本中的 main 函数
func (v *WrappedVm) Run(args []interface{}) (interface{}, error) {
	return v.RunFunc("main", args)
}

func (v *WrappedVm) RunFunc(funcName string, args []interface{}) (result interface{}, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			if errTemp, ok := recoverErr.(error); ok {
				err = errors.New(fmt.Sprintf("function %s run failed - %s", funcName, errTemp.Error()))
			} else {
				err = errors.New(fmt.Sprintf("function %s run failed", funcName))
			}
		}
	}()
	if args == nil {
		args = []interface{}{"undefined"} // 必须填充一个参数，否则编译报错。goja 的问题
	}
	mainFunc, err := v.mustFindFunc(funcName)
	if err != nil {
		return "", go_error.WithStack(err)
	}

	return mainFunc(args), nil
}

func (v *WrappedVm) mustFindFunc(funcName string) (result MainFuncType, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			if tempErr, ok := recoverErr.(error); ok && strings.Contains(tempErr.Error(), "invalid memory address or nil pointer dereference") {
				err = go_error.WithStack(errors.New("function not be found"))
				return
			}
		}
	}()
	var mainFunc MainFuncType
	err = v.Vm.ExportTo(v.Vm.Get(funcName), &mainFunc)
	if err != nil {
		return nil, go_error.WithStack(err)
	}
	return mainFunc, nil
}
