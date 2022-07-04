package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// ENV_NAME 环境变量名称
const ENV_NAME = "XW_DAEMON_IDX"

// 运行时调用background的次数
var runIdx int = 0

// Daemon 守护进程
type Daemon struct {
	MaxCount    int   //循环重启最大次数, 若为0则无限重启
	MaxError    int   //连续启动失败或异常退出的最大次数, 超过此数, 守护进程退出, 不再重启子进程
	MinExitTime int64 //子进程正常退出的最小时间(秒). 小于此时间则认为是异常退出
}

// Background 把本身程序转化为后台运行(启动一个子进程, 然后自己退出)
// logFile 若不为空,子程序的标准输出和错误输出将记入此文件
// isExit  启动子加进程后是否直接退出主程序, 若为false, 主程序返回*os.Process, 子程序返回 nil. 需自行判断处理
func Background(isExit bool) (*exec.Cmd, error) {
	// 判断子进程还是父进程
	runIdx++
	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	if runIdx <= envIdx { //子进程, 退出
		return nil, nil
	}

	// 设置子进程环境变量
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))

	// 启动子进程
	cmd, err := startProc(os.Args, env)
	if err != nil {
		return nil, err
	}

	if isExit {
		os.Exit(0)
	}

	return cmd, nil
}

// NewDaemon 创建守护进程对象
func NewDaemon() *Daemon {
	return &Daemon{
		MaxCount:    0,
		MaxError:    3,
		MinExitTime: 10,
	}
}

// Run 启动后台守护进程
func (d *Daemon) Run() {
	// 启动一个守护进程后退出
	_, err := Background(true)
	if err != nil {
		panic(err)
	}

	// 守护进程启动一个子进程, 并循环监视
	var t int64
	count := 1
	errNum := 0
	for {
		//daemon 信息描述
		if errNum > d.MaxError {
			os.Exit(1)
		}
		if d.MaxCount > 0 && count > d.MaxCount {
			os.Exit(0)
		}
		count++

		t = time.Now().Unix() //启动时间戳
		cmd, err := Background(false)
		if err != nil { //启动失败
			errNum++
			continue
		}

		//子进程,
		if cmd == nil {
			break
		}

		// 父进程: 等待子进程退出
		err = cmd.Wait()
		dat := time.Now().Unix() - t //子进程运行秒数
		if dat < d.MinExitTime {     //异常退出
			errNum++
		} else { //正常退出
			errNum = 0
		}
	}
}

// 启动进程
func startProc(args, env []string) (*exec.Cmd, error) {
	cmd := &exec.Cmd{
		Path:        args[0],
		Args:        args,
		Env:         env,
		SysProcAttr: NewSysProcAttr(),
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
