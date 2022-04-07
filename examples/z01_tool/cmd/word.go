package cmd

import (
	"log"
	"strings"

	"github.com/zhangdapeng520/zdpgo_cmd/examples/z01_tool/internal/word"
	"github.com/zhangdapeng520/zdpgo_cmd/libs/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var str string                    // 字符串
var mode int8                     // 模式
var desc = strings.Join([]string{ // 描述
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

func init() {
	// 解析单词内容
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")

	// 解析模式
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

// 创建命令
var wordCmd = &cobra.Command{
	Use:     "word",   // 子命令
	Short:   "单词格式转换", // 帮助描述
	Long:    desc,     // 详细描述
	Example: "z01_tool.exe word -m 1 -s hello_word",
	// 核心实现：命令的内容
	Run: func(cmd *cobra.Command, args []string) {
		// 内容
		var content string

		// 模式
		switch mode {
		case ModeUpper: // 大写
			content = word.ToUpper(str)
		case ModeLower: // 小写
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase: // 下划线转大驼峰
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase: // 下划线转小驼峰
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore: // 驼峰转下划线
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		// 将结果输出到控制台
		log.Printf("输出结果: %s", content)
	},
}
