package zdpgo_cmd

import (
	"github.com/zhangdapeng520/zdpgo_cmd/daemon"
	"os"
	"strconv"
)

func (c *Cmd) RunWithBackGround(runFunc func()) error {
	// 启动一个子进程后主程序退出
	_, _ = daemon.Background(Log.Config.LogFilePath, true)

	// 如果.env不存在，则创建
	_, err := os.Stat(c.Config.EnvFileName)
	if err != nil {
		_, err = os.Create(c.Config.EnvFileName)
	}

	// 将PID写入到环境变量
	err = Env.Load(c.Config.EnvFileName)
	if err != nil {
		Log.Error(err.Error())
		return err
	}
	var envMap = make(map[string]string)
	if fileMap, ok := Env.FileEnvMap[c.Config.EnvFileName]; ok {
		envMap = fileMap
	}
	envMap[c.Config.PidName] = strconv.Itoa(os.Getpid())

	// 写入环境变量
	err = Env.WriteNew(c.Config.EnvFileName, envMap)
	if err != nil {
		Log.Error(err.Error())
		return err
	}

	// 启动服务
	runFunc()

	return nil
}

// ExitBackground 退出后台任务
func (c *Cmd) ExitBackground() error {
	// 如果.env不存在，则创建
	_, err := os.Stat(c.Config.EnvFileName)
	if err != nil {
		_, err = os.Create(c.Config.EnvFileName)
	}

	// 加载环境变量
	err = Env.Load(c.Config.EnvFileName)
	if err != nil {
		Log.Error(err.Error())
		return err
	}

	// 获取进程号
	pid := Env.Get(c.Config.PidName)
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		Log.Error(err.Error())
		return err
	}

	// 关闭进程
	result := Shell.Kill(pidInt)
	Log.Debug("close process success", "result", result)
	return nil
}
