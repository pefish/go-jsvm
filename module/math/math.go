package math

import (
	"github.com/pefish/go-jsvm/module"
	"math"
)

const ModuleName = "Math"

type Math struct {
	vm module.IWrappedVm
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

func (c *Math) Max(datas []float64) float64 {
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

func (c *Math) Min(datas []float64) float64 {
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

func NewMathModule(vm module.IWrappedVm) *Math {
	return &Math{
		vm: vm,
	}
}
