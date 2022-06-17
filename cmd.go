package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_env"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_shell"
)

var (
	Log   *zdpgo_log.Log
	Shell = zdpgo_shell.New()
	Env   = zdpgo_env.New()
)

type Cmd struct {
	Config *Config
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
	Log = log
	return c
}
