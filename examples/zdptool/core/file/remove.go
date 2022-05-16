package file

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd/cobra"
)

/*
@Time : 2022/5/6 11:50
@Author : 张大鹏
@File : rename.go
@Software: Goland2021.3.1
@Description: 文件移除相关
*/

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "文件移除相关命令",
	Long:  "文件移除相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		// 移除文件夹的后缀
		if suffix {
			err := file.RemoveDirFilesSuffix(removeDirPath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("移除文件夹中所有文件的后缀成功")
		}
	},
}

var (
	suffix        bool
	removeDirPath string
)

func init() {
	removeCmd.Flags().BoolVarP(&suffix, "suffix", "s", false, `是否移除文件夹所有文件后缀`)
	removeCmd.Flags().StringVarP(&removeDirPath, "path", "p", "./", `文件夹目录`)
}
