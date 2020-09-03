package HTTP

import (
	"myCPforGo/Model"
)

//GetGameData 获取Game数据
type IGetGameData interface {
	/*
		返回某一年的数据
		@strYear 年份
		return []Model.GameAllBasic
	*/
	GetGameDataForYear() []Model.GameAllBasic
}
