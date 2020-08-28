package Test

import (
	"myCPforGo/Business/CPHttp/ImpMethod"
	"myCPforGo/Interface/HTTP"
	"testing"
)

func TestA(t *testing.T) {
	var httpMode HTTP.IGetGameData
	httpMode = ImpMethod.GetGameDataOne{conCond: HTTP.ConditionParams{Iyear: "2019"}}
	httpMode.GetGameDataForYear()

}
