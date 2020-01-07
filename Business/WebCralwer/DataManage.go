package WebCralwer

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

//Calculate_ScoringRate 计算进球率
//team:名称 number:场次
//return float64
func Calculate_ScoringRate(team string, number int) float64 {
	//1.根据team找到number场
	results := SearchForGame(team)
	if len(results) < number {
		number = len(results)
	}
	var sNumbers int
	for i := 0; i < number; i++ {
		//计算主客场数
		if results[i]["GguestName"] == team {
			sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], true)
		} else {
			sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], false)
		}
	}

	d1, _ := decimal.NewFromFloat(float64(number)).Div(decimal.NewFromFloat(float64(sNumbers))).Float64()
	return d1
}
func resolveSources(sources string, isHome bool) int {
	var scoring int
	if isHome {
		scoring, _ = strconv.Atoi(strings.Split(sources, "-")[0])
		return scoring

	}
	scoring, _ = strconv.Atoi(strings.Split(sources, "-")[1])
	return scoring
}
