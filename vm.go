package vm

import (
	"fmt"
	"io"
	"os"
	"strings"

	raw_errors "errors"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/buffer"
	"github.com/dop251/goja_nodejs/process"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/url"
	"github.com/pefish/go-jsvm/module/console"
	"github.com/pefish/go-jsvm/module/http"
	"github.com/pefish/go-jsvm/module/math"
	"github.com/pefish/go-jsvm/module/regex"
	"github.com/pefish/go-jsvm/module/time"
	go_logger "github.com/pefish/go-logger"
	"github.com/pkg/errors"
)

type WrappedVm struct {
	Vm     *goja.Runtime
	script string
	logger go_logger.InterfaceLogger

	ConsoleModule *console.Console
}

type MainFuncType func([]interface{}) interface{}

func (v *WrappedVm) SetLogger(logger go_logger.InterfaceLogger) *WrappedVm {
	v.logger = logger
	return v
}

func (v *WrappedVm) Logger() go_logger.InterfaceLogger {
	return v.logger
}

func NewVm(script string) *WrappedVm {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	wrappedVm := &WrappedVm{
		Vm:     vm,
		script: script,
		logger: go_logger.Logger,
	}
	wrappedVm.registerModules()
	return wrappedVm
}

func NewVmWithFile(jsFilename string) (*WrappedVm, error) {
	fileInfo, err := os.Stat(jsFilename)
	if err != nil {
		return nil, errors.Wrapf(err, "Illegal js file <%s>.", jsFilename)
	}
	if fileInfo.IsDir() || !fileInfo.Mode().IsRegular() {
		return nil, errors.Errorf("Illegal js file <%s>.", jsFilename)
	}
	f, err := os.Open(jsFilename)
	if err != nil {
		return nil, errors.Wrapf(err, "Illegal js file <%s>.", jsFilename)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "Illegal js file <%s>.", jsFilename)
	}
	return NewVm(string(content)), nil
}

// 注册预设的一些模块
func (v *WrappedVm) registerModules() {
	new(require.Registry).Enable(v.Vm)
	buffer.Enable(v.Vm)
	process.Enable(v.Vm)
	url.Enable(v.Vm)

	v.RegisterModule(regex.ModuleName, regex.NewRegexModule(v))

	mathModule := math.NewMathModule(v)
	v.RegisterModule(math.ModuleName, mathModule)
	v.RegisterModule(math.ModuleName1, mathModule)

	v.RegisterModule(time.ModuleName, time.NewTimeModule(v))

	v.RegisterModule(http.ModuleName, http.NewHttpModule(v))

	v.ConsoleModule = console.NewConsoleModule(v)
	v.RegisterModule(console.ModuleName, v.ConsoleModule)
	v.RegisterModule(console.ModuleName1, v.ConsoleModule)
}

func (v *WrappedVm) RegisterModule(moduleName string, module interface{}) error {
	err := v.Vm.Set(moduleName, module)
	if err != nil {
		return errors.Wrapf(err, "Register module <%s> error.", moduleName)
	}
	return nil
}

func (v *WrappedVm) ToValue(i interface{}) goja.Value {
	return v.Vm.ToValue(i)
}

func (v *WrappedVm) Panic(err error) {
	panic(
		v.ToValue(
			map[string]string{
				"message": fmt.Sprintf("%+v", err),
			},
		),
	)
}

func (v *WrappedVm) PanicWithMsg(msg string) {
	panic(
		v.ToValue(
			map[string]string{
				"message": fmt.Sprintf("%+v", errors.New(msg)),
			},
		),
	)
}

func (v *WrappedVm) RunMain(args []interface{}) (interface{}, error) {
	return v.RunFunc("main", args)
}

func (v *WrappedVm) Run() (goja.Value, error) {
	value, err := v.Vm.RunString(v.script)
	if err != nil {
		if strings.Contains(err.Error(), "Killed") {
			return nil, raw_errors.New("Killed")
		}
		return nil, errors.Wrapf(err, "Run script error.")
	}
	return value, nil
}

func (v *WrappedVm) Kill() {
	v.Vm.Interrupt("Killed")
}

func (v *WrappedVm) RunFunc(funcName string, args []interface{}) (result interface{}, err_ error) {
	defer func() {
		if err := recover(); err != nil {
			err_ = errors.Errorf("Function %s run failed - %s", funcName, err.(error).Error())
		}
	}()
	if args == nil {
		args = []interface{}{"undefined"} // 必须填充一个参数，否则编译报错。goja 的问题
	}
	mainFunc, err := v.findFunc(funcName)
	if err != nil {
		return "", errors.Errorf("Function %s run failed - %s", funcName, err.Error())
	}

	mainFuncResult := mainFunc(args) // panic when js throw
	return mainFuncResult, nil
}

func (v *WrappedVm) findFunc(funcName string) (result MainFuncType, err_ error) {
	defer func() {
		if err := recover(); err != nil {
			err_ = errors.Wrap(err.(error), fmt.Sprintf("Js function <%s> not be found", funcName))
		}
	}()
	var mainFunc MainFuncType
	jsFunc := v.Vm.Get(funcName) // panic when not found
	err := v.Vm.ExportTo(jsFunc, &mainFunc)
	if err != nil {
		return nil, errors.Wrapf(err, "Export js function <%s> error", funcName)
	}
	return mainFunc, nil
}
