package file

import (
	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_file"
)

// FileCmd 操作文件的相关命令
var FileCmd = &zdpgo_cmd.Command{
	Use:   "file",
	Short: "操作文件的相关命令",
	Long:  "操作文件的相关命令",
	Run:   func(cmd *zdpgo_cmd.Command, args []string) {},
}

var (
	file = zdpgo_file.New()
)

func init() {
	FileCmd.AddCommand(renameCmd)
	FileCmd.AddCommand(removeCmd)
}
