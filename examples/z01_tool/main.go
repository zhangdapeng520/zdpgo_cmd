package main

import (
	"log"

	"github.com/zhangdapeng520/zdpgo_cmd/examples/z01_tool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
