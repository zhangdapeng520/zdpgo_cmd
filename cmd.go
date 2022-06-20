package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_env"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_shell"
)

type Cmd struct {
	Config *Config
	Log    *zdpgo_log.Log
	Shell  *zdpgo_shell.Shell
	Env    *zdpgo_env.Env
}

func New(log *zdpgo_log.Log) *Cmd {
	return NewWithConfig(&Config{}, log)
}

func NewWithConfig(config *Config, log *zdpgo_log.Log) *Cmd {
	c := &Cmd{}
	if config.PidName == "" {
		config.PidName = "zdpgo_cmd_background_pid"
	}
	if config.EnvFileName == "" {
		config.EnvFileName = ".env"
	}
	c.Config = config
	c.Log = log
	c.Shell = zdpgo_shell.New()
	c.Env = zdpgo_env.New()
	return c
}
