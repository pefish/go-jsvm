package module

import (
	"github.com/dlclark/regexp2"
)

type Regex struct {
	vm IWrappedVm
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

func NewRegexModule(vm IWrappedVm) *Regex {
	return &Regex{
		vm: vm,
	}
}
