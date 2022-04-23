package random

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
	"github.com/zhangdapeng520/zdpgo_random"
)

// RandomCmd 随机数据子命令，子命令的入口
var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "生成随机数据",
	Long:  "生成随机数据",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var (
	length int
	result string
	random = zdpgo_random.New(zdpgo_random.RandomConfig{})
)

// 生成随机字符串
var strCmd = &cobra.Command{
	Use:   "str",
	Short: "生成随机的字符串",
	Long:  "生成随机的字符串",
	Run: func(cmd *cobra.Command, args []string) {
		result = random.Str.Str(length)
		fmt.Println(result)
	},
}

func init() {
	RandomCmd.AddCommand(strCmd)

	// 获取字符串的长度
	strCmd.Flags().IntVarP(&length, "length", "l", 16, `要生成的随机字符串的长度，需要传入一个整数`)
}
