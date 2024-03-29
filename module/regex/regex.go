package regex

import (
	"github.com/dlclark/regexp2"
	"github.com/pefish/go-jsvm/module"
)

const ModuleName = "regex"

type Regex struct {
	vm module.IWrappedVm
}

func (r *Regex) Match(pattern string, targetStr string) bool {
	reg, err := regexp2.Compile(pattern, 0)
	if err != nil {
		panic(r.vm.ToValue(err))
	}
	bool_, err := reg.MatchString(targetStr)
	if err != nil {
		panic(r.vm.ToValue(err))
	}
	return bool_
}

func NewRegexModule(vm module.IWrappedVm) *Regex {
	return &Regex{
		vm: vm,
	}
}
