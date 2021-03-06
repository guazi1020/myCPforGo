/**
* 对比赛的数据的操作
*
*
**/
package WebCralwer

import (
	_ "fmt"
	"math"
	"myCPforGo/Com/baseMethod"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	E = 2.718
)

//Probabiltiy_ScoringRate 预测进球率
//team:球队,exceptGlobals 预测进球数,lastNumber 最近几场,isHome 是否是主场(0,全部;1,主场；2,客场),league 赛制名称
func Probability_ScoringRate(team string, exceptGlobals int, lastNumber int, isHome int, league ...string) float64 {
	//公式：P（X）=（M^X/X!)*e^(-M)；P (0) = e^(-M)
	//M为球队场均进球数
	//X为期望进球值
	//e为常实数2.718
	//场均进球
	_lastNumber := 1
	var m float64
	if lastNumber > 0 {
		_lastNumber = lastNumber
	}
	if isHome == 0 {

	}
	//场均进球数
	switch isHome {
	case 1:
		m, _ = strconv.ParseFloat(baseMethod.ChangeNumber(Calculate_AveGlobal(team, _lastNumber, 1, league...), 3), 64)
	case 2:
		m, _ = strconv.ParseFloat(baseMethod.ChangeNumber(Calculate_AveGlobal(team, _lastNumber, 2, league...), 3), 64)
	default:
		m, _ = strconv.ParseFloat(baseMethod.ChangeNumber(Calculate_AveGlobal(team, _lastNumber, 0, league...), 3), 64)
	}
	//fmt.Println("m", m)
	avgGlobals, _ := decimal.NewFromFloat(float64(baseMethod.CountMultiplying(m, exceptGlobals))).Div(decimal.NewFromFloat(float64(baseMethod.CountFactorial(exceptGlobals)))).Float64()
	//fmt.Println("M^X/X!=", avgGlobals)
	//e^(-M)
	avgGlobals = avgGlobals * math.Pow(E, -m)
	avgGlobals, _ = strconv.ParseFloat(baseMethod.ChangeNumber(avgGlobals, 5), 64)
	return avgGlobals
}

//Calculate_sumGlobal 计算总进球数
//team:球队 count:轮数 Ishome:0,all 1,主场 2,客场
func Calculate_sumGlobal(team string, count int, Ishomes int, league ...string) int {
	var results map[int]map[string]string
	var sNumbers int
	switch Ishomes {
	default:
		break
	case 0:
		results = SearchForGame(team, count, 0, league...)
		for _, result := range results {
			if result["GAhomeName"] == team {
				sNumbers = sNumbers + resolveSources(result["GAresultScore"], true)
			} else {
				sNumbers = sNumbers + resolveSources(result["GAresultScore"], false)
			}
		}
	case 1:
		results = SearchForGame(team, count, 1, league...)
		for i := 0; i < len(results); i++ {
			sNumbers = sNumbers + resolveSources(results[i]["GAresultScore"], true)
		}
	case 2:
		results = SearchForGame(team, count, 2, league...)
		for i := 0; i < len(results); i++ {
			sNumbers = sNumbers + resolveSources(results[i]["GAresultScore"], false)
		}
	}

	return sNumbers
}

//Calulate_AveGlobal 平均进球数
//parame team 球队 tnumbers 轮数,ishomes 0,全部 1主场 2,客场
func Calculate_AveGlobal(team string, tnumbers int, ishomes int, league ...string) float64 {
	//根据team找到tnumbersshu
	var globals int //总进球数
	globals = Calculate_sumGlobal(team, tnumbers, ishomes, league...)

	avgGlobals, _ := decimal.NewFromFloat(float64(globals)).Div(decimal.NewFromFloat(float64(tnumbers))).Float64()
	avgGlobals, _ = strconv.ParseFloat(baseMethod.ChangeNumber(avgGlobals, 5), 64)
	return avgGlobals
}

//resovleSources 比分拆解
//sources 比分 ishome 是否主队(true主队,false客队)
func resolveSources(sources string, isHome bool) int {
	var scoring int
	if isHome {
		scoring, _ = strconv.Atoi(strings.Split(sources, "-")[0])
		return scoring

	}
	scoring, _ = strconv.Atoi(strings.Split(sources, "-")[1])
	return scoring
}

//Calculate_E 计算E值
//E=P*SP
func Calculate_E(D_value float64, main_sp float64) float64 {
	//E=P*main_SP(主)
	//P:(45%+0.53%*D_value)
	p := 0.45 + 0.0053*D_value
	return p * main_sp
}
