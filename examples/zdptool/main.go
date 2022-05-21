package main

import (
	"log"

	"zdptool/cmd"
)

func main() {
	// 执行CMD命令
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute 执行CMD命令失败: %v", err)
	}
}
