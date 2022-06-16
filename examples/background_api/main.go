package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_api"
	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_log"
)

var (
	rootCmd    = &zdpgo_cmd.Command{}
	background bool // 是否后台运行
	exit       bool // 是否退出
	debug      bool
	port       int
	Log        *zdpgo_log.Log
	Cmd        = zdpgo_cmd.New()
	Api        *zdpgo_api.Api
)

func init() {
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, `是否为Debug调试模式`)
	rootCmd.Flags().IntVarP(&port, "port", "p", 8888, `服务启动端口号`)
	rootCmd.Flags().BoolVarP(&background, "background", "b", false, `是否在后台运行服务`)
	rootCmd.Flags().BoolVarP(&exit, "exit", "e", false, `是否退出后台运行服务`)
}

func main() {
	// 创建一个根cmd对象
	rootCmd.Run = func(cmd *zdpgo_cmd.Command, args []string) {
		fmt.Println("args", "debug", debug, "port", port)
		Log = zdpgo_log.NewWithDebug(debug, "log.log")
		Api = zdpgo_api.NewWithConfig(&zdpgo_api.Config{
			Debug: debug,
			Port:  port,
		}, Log)

		Api.Post("/aes", func(ctx *zdpgo_api.Context) {
			// 解析json数据
			var jsonData struct {
				Username string `json:"username"`
				Age      int    `json:"age"`
			}
			err := ctx.GetAesTextBodyToJson(&jsonData)
			if err != nil {
				panic(err)
			}

			// 加密响应数据
			response := &zdpgo_api.Response{
				Code:   10000,
				Msg:    "success",
				Status: true,
				Data:   jsonData,
			}
			ctx.ResponseAesStr(response)
		})

		if background {
			Cmd.RunWithBackGround(func() {
				Api.Run()
			})
		} else {
			Api.Run()
		}

		// 处理退出
		if exit {
			err := Cmd.ExitBackground()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("退出后台服务成功")
		}
	}

	// 执行根rootCmd的命令
	rootCmd.Execute()
}
