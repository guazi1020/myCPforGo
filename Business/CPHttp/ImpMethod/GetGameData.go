package ImpMethod

import (
	"fmt"
	"myCPforGo/Model"
)

type GetGameDataOne struct {
	// conCond HTTP.ConditionParams
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear() []Model.GameAllBasic {

	var games []Model.GameAllBasic
	fmt.Println("abcdefg")
	return games
}
