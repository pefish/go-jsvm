package module

import (
	"github.com/dlclark/regexp2"
)

type Regex struct {
	Match func(pattern string, targetStr string) bool `json:"match"`
}

func RegisterRegexModule(v IWrappedVm) error {
	return v.RegisterModule("regex", &Regex{
		Match: func(pattern string, targetStr string) bool {
			reg, err := regexp2.Compile(pattern, 0)
			if err != nil {
				panic(v.ToValue(err))
			}
			bool_, err := reg.MatchString(targetStr)
			if err != nil {
				panic(v.ToValue(err))
			}
			return bool_
		},
	})
}
