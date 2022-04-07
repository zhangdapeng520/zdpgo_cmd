package cmd

import (
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
)

// 创建一个根cmd对象
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

// Execute 执行CMD命令
func Execute() error {
	// 执行根rootCmd的命令
	return rootCmd.Execute()
}

func init() {
	// 向根命令添加word子命令
	rootCmd.AddCommand(wordCmd)

	// 向根命令添加time子命令
	rootCmd.AddCommand(timeCmd)

	// 向根命令添加json子命令
	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(sqlCmd)
}
