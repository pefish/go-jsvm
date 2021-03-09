package vm

import (
	"github.com/pefish/go-test-assert"
	"strings"
	"testing"
)

func TestNewVmAndLoad(t *testing.T) {
	vm, err := NewVmAndLoad(`
console.log("123")
`)
	test.Equal(t, nil, err)
	_, err = vm.Run(nil)
	test.Equal(t, true, err != nil && err.Error() == "function not be found")

	vm1, err := NewVmAndLoad(`
function main(args) {
console.log(args)
return args
}
`)
	test.Equal(t, nil, err)
	result, err := vm1.Run([]interface{}{"111","222", 123})
	test.Equal(t, nil, err)
	realResult, ok := result.([]interface{})
	test.Equal(t, true, ok)
	test.Equal(t, "111", realResult[0].(string))
	test.Equal(t, "222", realResult[1].(string))
	test.Equal(t, 123, realResult[2].(int))

	vm2, err := NewVmAndLoad(`
function main(args) {
	return regex.match(args[0], args[1])
}
`)
	test.Equal(t, nil, err)
	_, err = vm2.Run([]interface{}{"*","222"})
	test.Equal(t, true, err != nil && strings.Contains(err.Error(), "function main run failed - error parsing regexp: missing argument to repetition operator"))

	result21, err := vm2.Run([]interface{}{".*","222"})
	test.Equal(t, nil, err)
	realResult21, ok := result21.(bool)
	test.Equal(t, true, ok)
	test.Equal(t, true, realResult21)
}
