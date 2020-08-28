package ImpMethod

import (
	"myCPforGo/Interface/HTTP"
	"myCPforGo/Model"
)

type GetGameDataOne struct {
	conCond HTTP.ConditionParams
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear() []Model.GameAllBasic {

	var games []Model.GameAllBasic

	return games
}
