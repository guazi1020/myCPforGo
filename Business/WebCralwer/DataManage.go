package WebCralwer

import (
	"fmt"
	"myCPforGo/Com/baseMethod"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	E = 2.718
)

//Probabiltiy_ScoringRate 预测进球率
//team:球队,exceptGlobals 预测进球数
func Probability_ScoringRate(team string, exceptGlobals int) float64 {
	//公式：P（X）=（M^X/X!)*e^(-M)；P (0) = e^(-M)
	//M为球队场均进球数
	//X为期望进球值
	//e为常实数2.718
	//场均进球

	m := Calculate_AveGlobal(team, 1)
	x := baseMethod.CountMultiplying(m, exceptGlobals)
	fmt.Println(x)
	y := baseMethod.CountFactorial(exceptGlobals)
	fmt.Println(y)
	//	fmt.Println(float64(baseMethod.CountMultiplying(m, exceptGlobals)))
	avgGlobals, _ := decimal.NewFromFloat(float64(baseMethod.CountMultiplying(m, exceptGlobals))).Div(decimal.NewFromFloat(float64(baseMethod.CountFactorial(exceptGlobals)))).Float64()
	avgGlobals = avgGlobals * baseMethod.CountMultiplying(E, int(-m))
	//	avgGlobals := baseMethod.CountMultiplying(E, -exceptGlobals)
	avgGlobals, _ = strconv.ParseFloat(baseMethod.ChangeNumber(avgGlobals, 5), 64)
	return avgGlobals
}

//Calculate_sumGlobal 计算总进球数
//team:球队 count:轮数 Ishome:是否是主客场,ture主场,guest客场,没有就是不管主客场
func Calculate_sumGlobal(team string, count int, Ishomes ...bool) int {
	var results map[int]map[string]string
	var sNumbers int

	if len(Ishomes) > 0 { //确定了主客场
		for _, ishome := range Ishomes {
			results = SearchForGame(team, count, ishome)
			switch ishome {
			case false:
				for i := 0; i < len(results); i++ {
					sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], false)
				}
			case true:
				for i := 0; i < len(results); i++ {
					sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], true)
				}
			default:
				break
			}
		}
	} else { //全部不确定主客场
		results = SearchForGame(team, count)
		for i := 0; i < len(results); i++ {
			if results[i]["GguestName"] == team { //主场比分
				sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], false)
			} else { //客场比分
				sNumbers = sNumbers + resolveSources(results[i]["GresultScore"], true)
			}

		}
	}

	return sNumbers
}

//Calulate_AveGlobal 平均进球数
//parame team 球队 tnumbers 轮数,ishomes 是否是主场,true/false 不填写全部
func Calculate_AveGlobal(team string, tnumbers int, ishomes ...bool) float64 {
	//根据team找到tnumbersshu
	var globals int //总进球数

	if len(ishomes) > 0 {
		for _, ishome := range ishomes {
			globals = Calculate_sumGlobal(team, tnumbers, ishome)
		}
	} else {
		globals = Calculate_sumGlobal(team, tnumbers)
	}
	avgGlobals, _ := decimal.NewFromFloat(float64(globals)).Div(decimal.NewFromFloat(float64(tnumbers))).Float64()
	avgGlobals, _ = strconv.ParseFloat(baseMethod.ChangeNumber(avgGlobals, 5), 64)
	return avgGlobals
}

//Calculate_ScoringRate 计算进球率
//team:名称 number:场次
//return float64
func Calculate_ScoringRate(team string, number int) float64 {
	//1.根据team找到number场

	var sNumbers int
	sNumbers = Calculate_sumGlobal(team, number, false)
	fmt.Println(sNumbers)
	d1, _ := decimal.NewFromFloat(float64(number)).Div(decimal.NewFromFloat(float64(sNumbers))).Float64()
	d1, _ = strconv.ParseFloat(baseMethod.ChangeNumber(d1, 5), 64)
	return d1
	//return sNumbers
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
