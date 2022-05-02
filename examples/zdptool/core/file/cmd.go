package file

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
	"github.com/zhangdapeng520/zdpgo_file"
)

// FileCmd 操作文件的相关命令
var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "操作文件的相关命令",
	Long:  "操作文件的相关命令",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var (
	dirPath string
	oldStr  string
	newStr  string
	file    = zdpgo_file.New()
)

// 生成随机字符串
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "重命名文件",
	Long:  "重命名文件",
	Run: func(cmd *cobra.Command, args []string) {
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

func init() {
	FileCmd.AddCommand(renameCmd)

	// 指定文件夹
	renameCmd.Flags().StringVarP(&dirPath, "dir", "d", "./", `指定要重命名文件的目录`)
	renameCmd.Flags().StringVarP(&oldStr, "oldStr", "o", "", `旧的字符串`)
	renameCmd.Flags().StringVarP(&newStr, "newStr", "n", "", `新的字符串`)
}
