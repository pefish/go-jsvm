package test_module

import (
	"fmt"
	"github.com/pefish/go-jsvm/module"
	"github.com/pkg/errors"
)

const ModuleName = "test_module"

type TestModule struct {
	vm module.IWrappedVm
}

func (t *TestModule) Test() {
	fmt.Println("test")
}

func (t *TestModule) TestPanic() {
	//t.vm.PanicWithMsg("test panic")
	t.vm.Panic(errors.New("test panic"))
}

type TestObj struct {
	A string `json:"a"`
}

func (t *TestModule) TestPtr() *TestObj {
	return &TestObj{A: "test"}
}

func (t *TestModule) TestNull() *TestObj {
	return nil
}

func NewTestModuleModule(vm module.IWrappedVm) *TestModule {
	return &TestModule{
		vm: vm,
	}
}
