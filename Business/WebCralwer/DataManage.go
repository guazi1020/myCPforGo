package WebCralwer

import (
	"myCPforGo/Com/baseMethod"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

//Probabiltiy_ScoringRate 预测进球率
//team:球队
func Probability_ScoringRate(team string, exceptGlobals int) float64 {
	//公式：P（X）=（M^X/X!)*e^(-M)；P (0) = e^(-M)
	//M为球队场均进球数
	//X为期望进球值
	//e为常实数2.718

	var rate float64
	return rate
}

//Calulate_AveGlobal 平均进球数
//parame team 球队
func Calculate_AveGlobal(team string, tnumbers int, pnumbers int) float64 {
	//根据team找到tnumbersshu
	var global float64
	// results := SearchForGame(team)
	// if len(results) < number {
	// 	number = len(results)
	// }
	// var sNumbers int
	// for i := 0; i < number; i++ {
	// 	//计算主客场数
	// 	if results[i]["GguestName"] == team {
	// 		sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], true)
	// 	} else {
	// 		sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], false)
	// 	}
	// }

	return global
}

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
	d1, _ = strconv.ParseFloat(baseMethod.ChangeNumber(d1, 5), 64)
	return d1
}

//resovleSources 比分拆解
//sources 比分 ishome 是否主队
func resolveSources(sources string, isHome bool) int {
	var scoring int
	if isHome {
		scoring, _ = strconv.Atoi(strings.Split(sources, "-")[0])
		return scoring

	}
	scoring, _ = strconv.Atoi(strings.Split(sources, "-")[1])
	return scoring
}
