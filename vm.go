package vm

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/buffer"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/process"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/url"
	"github.com/pefish/go-jsvm/module/math"
	"github.com/pefish/go-jsvm/module/regex"
	go_logger "github.com/pefish/go-logger"
	"github.com/pkg/errors"
	"io"
	"os"
)

type WrappedVm struct {
	Vm     *goja.Runtime
	script string
	logger go_logger.InterfaceLogger
}

type MainFuncType func([]interface{}) interface{}

func (v *WrappedVm) SetLogger(logger go_logger.InterfaceLogger) *WrappedVm {
	v.logger = logger
	return v
}

func (v *WrappedVm) Logger() go_logger.InterfaceLogger {
	return v.logger
}

func NewVm(script string) (*WrappedVm, error) {

	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	wrappedVm := &WrappedVm{
		Vm:     vm,
		script: script,
		logger: go_logger.Logger,
	}
	err := wrappedVm.registerModules()
	if err != nil {
		return nil, err
	}
	return wrappedVm, nil
}

func NewVmWithFile(jsFilename string) (*WrappedVm, error) {
	fileInfo, err := os.Stat(jsFilename)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		return nil, errors.New("illegal js file")
	}
	f, err := os.Open(jsFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	vm, err := NewVm(string(content))
	if err != nil {
		return nil, err
	}
	return vm, nil
}

// 注册预设的一些模块
func (v *WrappedVm) registerModules() error {
	new(require.Registry).Enable(v.Vm)
	console.Enable(v.Vm)
	buffer.Enable(v.Vm)
	process.Enable(v.Vm)
	url.Enable(v.Vm)

	err := v.RegisterModule(regex.ModuleName, regex.NewRegexModule(v))
	if err != nil {
		return err
	}

	err = v.RegisterModule(math.ModuleName, math.NewMathModule(v))
	if err != nil {
		return err
	}

	return nil
}

func (v *WrappedVm) RegisterModule(moduleName string, module interface{}) error {
	err := v.Vm.Set(moduleName, module)
	if err != nil {
		return err
	}
	return nil
}

func (v *WrappedVm) ToValue(i interface{}) goja.Value {
	return v.Vm.ToValue(i)
}

func (v *WrappedVm) Panic(err error) {
	panic(v.ToValue(fmt.Sprintf("%+v", err)))
}

func (v *WrappedVm) PanicWithMsg(msg string) {
	panic(v.ToValue(fmt.Sprintf("%+v", errors.New(msg))))
}

func (v *WrappedVm) RunMain(args []interface{}) (interface{}, error) {
	return v.RunFunc("main", args)
}

func (v *WrappedVm) Run() (goja.Value, error) {
	return v.Vm.RunString(v.script)
}

func (v *WrappedVm) RunFunc(funcName string, args []interface{}) (result interface{}, err_ error) {
	defer func() {
		if err := recover(); err != nil {
			err_ = errors.Errorf("function %s run failed - %s", funcName, err.(error).Error())
		}
	}()
	if args == nil {
		args = []interface{}{"undefined"} // 必须填充一个参数，否则编译报错。goja 的问题
	}
	mainFunc, err := v.findFunc(funcName)
	if err != nil {
		return "", errors.Errorf("function %s run failed - %s", funcName, err.Error())
	}

	mainFuncResult := mainFunc(args) // panic when js throw
	return mainFuncResult, nil
}

func (v *WrappedVm) findFunc(funcName string) (result MainFuncType, err_ error) {
	defer func() {
		if err := recover(); err != nil {
			err_ = errors.Wrap(err.(error), fmt.Sprintf("js function <%s> not be found", funcName))
		}
	}()
	var mainFunc MainFuncType
	jsFunc := v.Vm.Get(funcName) // panic when not found
	err := v.Vm.ExportTo(jsFunc, &mainFunc)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("export js function <%s> error", funcName))
	}
	return mainFunc, nil
}
