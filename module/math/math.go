package math

import (
	"math"

	go_format "github.com/pefish/go-format"
	"github.com/pefish/go-jsvm/module"
	"github.com/pkg/errors"
)

const ModuleName = "Math"
const ModuleName1 = "math_go"

type Math struct {
	vm module.IWrappedVm
}

func NewMathModule(vm module.IWrappedVm) *Math {
	return &Math{
		vm: vm,
	}
}

func (c *Math) Abs(a float64) float64 {
	return math.Abs(a)
}

func (c *Math) Average(values []float64) float64 {
	sum := 0.0
	count := len(values)

	// 计算总和
	for _, v := range values {
		sum += v
	}

	return sum / float64(count)
}

func (c *Math) AverageObjByKey(objs []map[string]interface{}, key string) float64 {
	sum := 0.0
	count := len(objs)

	// 计算总和
	for _, v := range objs {
		f, err := go_format.FormatInstance.ToFloat64(v[key])
		if err != nil {
			c.vm.Panic(errors.Wrap(err, ""))
		}
		sum += f
	}

	return sum / float64(count)
}

func (c *Math) Max(datas ...float64) float64 {
	result := datas[0]
	for i := 1; i < len(datas); i++ {
		if math.IsNaN(result) {
			result = datas[i]
			continue
		}
		if !math.IsNaN(datas[i]) && datas[i] > result {
			result = datas[i]
		}
	}
	return result
}

func (c *Math) Min(datas ...float64) float64 {
	result := datas[0]
	for i := 1; i < len(datas); i++ {
		if math.IsNaN(result) {
			result = datas[i]
			continue
		}
		if !math.IsNaN(datas[i]) && datas[i] < result {
			result = datas[i]
		}
	}
	return result
}
