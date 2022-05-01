package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_cmd/libs/daemon"
	"os"
	"strconv"
)

func (c *Cmd) RunWithBackGround(runFunc func()) error {
	// 启动一个子进程后主程序退出
	_, _ = daemon.Background(c.Config.LogFileName, true)

	// 将PID写入到环境变量
	err := c.Env.Load(c.Config.EnvFileName)
	if err != nil {
		c.Log.Error(err.Error())
		return err
	}
	var envMap = make(map[string]string)
	if fileMap, ok := c.Env.FileEnvMap[c.Config.EnvFileName]; ok {
		envMap = fileMap
	}
	envMap[c.Config.PidName] = strconv.Itoa(os.Getpid())

	// 写入环境变量
	err = c.Env.WriteNew(c.Config.EnvFileName, envMap)
	if err != nil {
		c.Log.Error(err.Error())
		return err
	}

	// 启动服务
	runFunc()

	return nil
}

func (c *Cmd) ExitBackground() error {
	err := c.Env.Load(".env")
	if err != nil {
		c.Log.Error(err.Error())
		return err
	}
	pid := c.Env.Get(c.Config.PidName)
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		c.Log.Error(err.Error())
		return err
	}
	c.Shell.Kill(pidInt)
	return nil
}
