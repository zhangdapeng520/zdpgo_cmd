package main

import (
	"log"

	"github.com/zhangdapeng520/zdpgo_cmd/examples/z01_tool/cmd"
)

func main() {
	// 执行CMD命令
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute 执行CMD命令失败: %v", err)
	}
}
