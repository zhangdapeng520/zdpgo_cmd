package zdpgo_cmd

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/zhangdapeng520/zdpgo_daemon"
)

func (c *Cmd) RunWithBackGround(runFunc func()) {
	// 启动一个子进程后主程序退出
	zdpgo_daemon.Background("", true)

	// 写入进程ID
	ioutil.WriteFile(".env", []byte(strconv.Itoa(os.Getpid())), os.ModePerm)

	// 启动服务
	runFunc()
}

// ExitBackground 退出后台任务
func (c *Cmd) ExitBackground() {
	// 关闭进程
	c.Shell.KillEnv()
}
