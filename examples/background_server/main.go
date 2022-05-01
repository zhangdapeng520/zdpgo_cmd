package main

import (
	"embed"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
	"html/template"
	"net/http"
)

//go:embed templates
var fs embed.FS

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 使用嵌入的文件系统，能够直接打包到二进制文件
	t, err := template.ParseFS(fs, "templates/index.html")

	if err != nil {
		panic(err)
	}

	// 解析模板
	t.Execute(w, nil)
}

var (
	rootCmd    = &cobra.Command{}
	background bool // 是否后台运行
	exit       bool // 是否退出
)

func init() {
	rootCmd.Flags().BoolVarP(&background, "background", "b", false, `是否在后台运行服务`)
	rootCmd.Flags().BoolVarP(&exit, "exit", "e", false, `是否退出后台运行服务`)
}

func main() {
	c := zdpgo_cmd.New()

	// 创建一个根cmd对象
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if background {
			c.RunWithBackGround(func() {
				server := &http.Server{Addr: ":8888"}
				http.HandleFunc("/", helloHandler)
				server.ListenAndServe()
			})
		} else {
			server := &http.Server{Addr: ":8888"}
			http.HandleFunc("/", helloHandler)
			server.ListenAndServe()
		}

		// 处理退出
		if exit {
			err := c.ExitBackground()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("退出后台服务成功")
		}
	}

	// 执行根rootCmd的命令
	rootCmd.Execute()
}
