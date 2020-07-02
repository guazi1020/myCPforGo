package Com

import (
	"regexp"
	"strings"
)

func RemoveBlank(str_source string) string {
	str_source = strings.Replace(str_source, " ", "", -1)
	// 去除换行符
	str_source = strings.Replace(str_source, "\n", "", -1)
	return str_source
}

//RemoveCharacter 替换字符
//str_source 原始字符串
//str 需要替换成为的字符
//stred 需要被替换的字符
func RemoveCharacter(str_source string, str string, stred string) string {
	return strings.Replace(str_source, stred, str, -1)
}

//JudeStringIsInt 判断字符是数字
//str 需要判断的字符串
//返回bool 判断是否正确
func JudgeStringIsInt(str string) bool {
	pattern := "\\d+" //正则判断是否是数字
	result, _ := regexp.MatchString(pattern, string(str))
	return result
}
