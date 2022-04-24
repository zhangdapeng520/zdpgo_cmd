package main

import (
	"embed"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd/libs/daemon"
	"html/template"
	"net/http"
	"os"
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

func main() {
	logFile := "daemon.log"

	fmt.Println("启动服务：http://localhost:8888")

	// 启动一个子进程后主程序退出
	daemon.Background(logFile, true)

	// 以下代码只有子程序会执行
	fmt.Println(os.Getpid(), "启动服务...")

	// 启动服务
	server := &http.Server{Addr: ":8888"}
	http.HandleFunc("/", helloHandler)
	server.ListenAndServe()

	// 服务退出
	fmt.Println(os.Getpid(), "服务退出")
}
