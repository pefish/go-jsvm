package main

import (
	"github.com/pefish/go-commander"
	"github.com/pefish/go-jsvm/cmd/main/command"
	"github.com/pefish/go-jsvm/version"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" 是一款使用运行 Javascript 代码的工具，祝你玩得开心。作者：pefish")
	commanderInstance.RegisterDefaultSubcommand(command.NewDefaultCommand())
	commanderInstance.DisableSubCommand()
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
