package module

type Console struct {
	Log func(data ...interface{}) `json:"log"`
}

func RegisterConsoleModule(v IWrappedVm) error {
	return v.RegisterModule("console", &Console{
		Log: func(data ...interface{}) {
			v.Logger().Info(data...)
		},
	})
}
