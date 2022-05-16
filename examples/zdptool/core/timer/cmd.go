package timer

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/zhangdapeng520/zdpgo_cmd/cobra"
)

// 时间子命令，子命令的入口
var TimeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// 当前时间子命令
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := getNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

// 计算的时间
var calculateTime string

// 间隔时间
var duration string

func init() {
	// 添加当前时间子命令
	TimeCmd.AddCommand(nowTimeCmd)

	// 添加计算时间子命令
	TimeCmd.AddCommand(calculateTimeCmd)

	// 获取要计算的时间
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)

	// 获取时间间隔
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}

// 计算时间子命令
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",

	// 核心方法，命令的实现
	Run: func(cmd *cobra.Command, args []string) {
		// 当前时间
		var currentTimer time.Time

		// 时间格式化
		var layout = "2006-01-02 15:04:05"

		// 要计算的时间为空
		if calculateTime == "" {
			currentTimer = getNowTime()
		} else {
			var err error
			// 用空格分割
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}

			// 解析时间
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		// 计算时间
		t, err := getCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		// 输出计算结果
		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}
