package vm

import (
	"strings"
	"testing"

	go_test_ "github.com/pefish/go-test"
)

func TestNewVm(t *testing.T) {
	vm, err := NewVm(`
console.log("123")
`)
	go_test_.Equal(t, nil, err)
	_, err = vm.Run()
	go_test_.Equal(t, true, err != nil && strings.Contains(err.Error(), "js function <main> not be found"))
}
