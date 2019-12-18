package Com

import "strings"

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
