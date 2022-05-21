package cmd

import (
	"github.com/zhangdapeng520/zdpgo_cmd"
	"zdptool/core/file"
	"zdptool/core/random"
	"zdptool/core/timer"
	"zdptool/core/word"
)

// 创建一个根cmd对象
var rootCmd = &zdpgo_cmd.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(word.WordCmd)     // 向根命令添加word子命令
	rootCmd.AddCommand(timer.TimeCmd)    // 向根命令添加time子命令
	rootCmd.AddCommand(random.RandomCmd) // 添加random随机数据命令
	rootCmd.AddCommand(file.FileCmd)
}

// Execute 执行CMD命令
func Execute() error {
	return rootCmd.Execute() // 执行根rootCmd的命令
}
