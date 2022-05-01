package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_env"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_shell"
)

type Cmd struct {
	Config *Config
	Shell  *zdpgo_shell.Shell
	Env    *zdpgo_env.Env
	Log    *zdpgo_log.Log
}

func New() *Cmd {
	return NewWithConfig(Config{})
}

func NewWithConfig(config Config) *Cmd {
	c := Cmd{}
	if config.PidName == "" {
		config.PidName = "zdpgo_cmd_background_pid"
	}
	if config.EnvFileName == "" {
		config.EnvFileName = ".env"
	}
	if config.LogFileName == "" {
		config.LogFileName = "logs/zdpgo/zdpgo_cmd.log"
	}
	c.Config = &config
	c.Shell = zdpgo_shell.New()
	c.Env = zdpgo_env.New()
	c.Log = zdpgo_log.NewWithConfig(zdpgo_log.Config{
		Debug:        config.Debug,
		OpenJsonLog:  true,
		OpenFileName: false,
		LogFilePath:  config.LogFileName,
	})
	return &c
}
