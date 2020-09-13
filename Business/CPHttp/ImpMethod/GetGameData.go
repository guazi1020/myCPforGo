package ImpMethod

import (
	"fmt"
	"myCPforGo/Business/DB"
	"myCPforGo/Model"
)

type GetGameDataOne struct {
	// conCond HTTP.ConditionParams
	Year string
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear() []Model.GameAllBasic {

	var games []Model.GameAllBasic
	var results map[int]map[string]string
	enable := DB.Case()
	strQuery := "select * from GameAllBasic where GADate>?"
	results = enable.Query(strQuery, getGameData.Year)
	fmt.Println(results)
	return games
}
