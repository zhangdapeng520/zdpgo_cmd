package word

import (
	"strings"
	"unicode"
)

// ToUpper 将单词转换为全大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 将单词转换为全小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转小驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	// 准备utf8字符数组
	var output []rune

	for i, r := range s {
		// 第一个单词小写
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		// 如果是大写，则追加一个下划线
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		// 追加单词的小写
		output = append(output, unicode.ToLower(r))
	}

	// 返回结果
	return string(output)
}
