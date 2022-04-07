package cmd

import (
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(sqlCmd)
}
