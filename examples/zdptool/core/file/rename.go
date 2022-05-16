package file

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd"
)

/*
@Time : 2022/5/6 11:50
@Author : 张大鹏
@File : rename.go
@Software: Goland2021.3.1
@Description: 重命名文件
*/

// 生成随机字符串
var renameCmd = &zdpgo_cmd.Command{
	Use:   "rename",
	Short: "重命名文件",
	Long:  "重命名文件",
	Run: func(cmd *zdpgo_cmd.Command, args []string) {
		if oldStr == "" {
			fmt.Println("要替换的字符串不能为空")
			return
		}
		err := file.ReplaceDirFilesName(dirPath, oldStr, newStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

var (
	dirPath string
	oldStr  string
	newStr  string
)

func init() {
	// 指定文件夹
	renameCmd.Flags().StringVarP(&dirPath, "dir", "d", "./", `指定要重命名文件的目录`)
	renameCmd.Flags().StringVarP(&oldStr, "oldStr", "o", "", `旧的字符串`)
	renameCmd.Flags().StringVarP(&newStr, "newStr", "n", "", `新的字符串`)
}
