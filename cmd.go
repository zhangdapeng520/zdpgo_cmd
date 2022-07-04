package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_env"
	"github.com/zhangdapeng520/zdpgo_shell"
)

type Cmd struct {
	Config *Config
	Shell  *zdpgo_shell.Shell
	Env    *zdpgo_env.Env
}

func New() *Cmd {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Cmd {
	c := &Cmd{}
	if config.PidName == "" {
		config.PidName = "zdpgo_cmd_background_pid"
	}
	if config.EnvFileName == "" {
		config.EnvFileName = ".env"
	}
	c.Config = config
	c.Shell = zdpgo_shell.New()
	c.Env = zdpgo_env.New()
	return c
}
