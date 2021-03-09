package module

import (
	"regexp"
)

type Regex struct {
	Match func(pattern string, targetStr string) bool `json:"match"`
}

func RegisterRegexModule(v IWrappedVm) error {
	return v.RegisterModule("regex", &Regex{
		Match: func(pattern string, targetStr string) bool {
			bool_, err := regexp.Match(pattern, []byte(targetStr))
			if err != nil {
				panic(v.ToValue(err))
			}
			return bool_
		},
	})
}
