package baseMethod

import (
	"strconv"
)

//CalculateGameResult 计算比分结果
//str 比分的string类
//返回为3/1/0
//例子： CalculateGameResult("1-0") 3
func CalculateGameResult(str string) string {

	//str = "2-2"

	for i, s := range str {
		if s == 45 {
			homeResource, _ := strconv.Atoi(str[:i])
			guestResource, _ := strconv.Atoi(str[i+1:])
			if homeResource > guestResource {
				return "3"
			}
			if homeResource < guestResource {
				return "0"
			}
			if homeResource == guestResource {
				return "1"
			}

		}

	}
	// fmt.Println(str[1])
	return "-1"
}
