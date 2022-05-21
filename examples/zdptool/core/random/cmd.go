package random

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd"
	"zdptool/g"
)

// RandomCmd 随机数据子命令，子命令的入口
var RandomCmd = &zdpgo_cmd.Command{
	Use:   "random",
	Short: "生成随机数据",
	Long:  "生成随机数据",
	Run:   func(cmd *zdpgo_cmd.Command, args []string) {},
}

var (
	length int
	result string
)

// 生成随机字符串
var strCmd = &zdpgo_cmd.Command{
	Use:   "str",
	Short: "生成随机的字符串",
	Long:  "生成随机的字符串",
	Run: func(cmd *zdpgo_cmd.Command, args []string) {
		result = g.Random.Str(length)
		fmt.Println(result)
	},
}

func init() {
	RandomCmd.AddCommand(strCmd)

	// 获取字符串的长度
	strCmd.Flags().IntVarP(&length, "length", "l", 16, `要生成的随机字符串的长度，需要传入一个整数`)
}
