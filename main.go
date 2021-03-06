package main

import (
	"fmt"
	_ "myCPforGo/Business/CPHttp"
	"myCPforGo/Business/WebCralwer"
	_ "myCPforGo/Com/baseMethod"
	"myCPforGo/Model"
	"reflect"
)

func main() {

	//CPHttp.StartHttp()
	/*
		更新一段时间的比赛信息
	*/
	var begindate, enddate string
	fmt.Print("开始日期:")
	fmt.Scanf("%s\n", &begindate)
	fmt.Print("结束日期:")
	fmt.Scanf("%s\n", &enddate)
	params := make(map[string]string)
	params["code"] = "all"
	params["ajax"] = "true"
	WebCralwer.SaveWebByDate(begindate, enddate, params)

	// //测算当前日期的比赛E和相关进球率预测
	// WebCralwer.GetEByDate(10)

	WebCralwer.GetEbyDateLiveStreaming(10, "2020-12-16")

}

func TestForReflect(igame interface{}) {
	game := reflect.ValueOf(igame)
	var igame1 Model.Game
	igame1 = game.Interface().(Model.Game)
	fmt.Println(igame1.GIsfinish)
}

//Equation 最终计算公式
func Equation() {
	//fmt.Println(Com.RemoveBlank("main, i am home,"))
	/**/
	params := make(map[string]string)
	params["code"] = "201"
	params["ajax"] = "true"
	//WebCralwer.SaveWebByDate("2016-01-01", "", params)
	//WebCralwer.MysqlDemo_Select()
	//fmt.Println(WebCralwer.Calculate_ScoringRate("切沃", 6))
	//fmt.Println(baseMethod.Compoundrate(152756, 0.0385, 24))
	//fmt.Println(WebCralwer.Calculate_AveGlobal("切沃", 3, true))
	//fmt.Println(baseMethod.CountFactorial(1))

	var team string    //球队名称
	var goals int      //进球数
	var num int        //几场比赛
	var ishome int     //主客场d
	var _ishome string //翻译临时主客场
	var league []string
	team = "格罗兹尼"
	goals = 0
	num = 8
	ishome = 1

	switch ishome {
	case 0:
		_ishome = "主客场"
	case 1:
		_ishome = "主场"
	case 2:
		_ishome = "客场"
	}
	league = append(league, "俄超")
	//	league = append(league, "欧洲杯")

	fmt.Printf("%s,进%d个球，范围为最近%d场%s，赛制为%s的情况下的可能性为：%f\n", team, goals, num, _ishome, league, WebCralwer.Probability_ScoringRate(team, goals, num, ishome, league...))

	team = "索契"
	goals = 1
	num = 8
	ishome = 2

	switch ishome {
	case 0:
		_ishome = "主客场"
	case 1:
		_ishome = "主场"
	case 2:
		_ishome = "客场"
	}
	fmt.Printf("%s,进%d个球，范围为最近%d场%s，赛制为%s的情况下的可能性为：%f\n", team, goals, num, _ishome, league, WebCralwer.Probability_ScoringRate(team, goals, num, ishome, league...))
	//fmt.Println(baseMethod.MyPow(4, 3))
	//fmt.Println(math.Pow(2.14, -1.23))
	//baseMethod.CountMultiplyingsqrt(2, 3.3)

	// var x int64
	// x = 7
	// var y int64
	// y = 5
	// fmt.Println(decimal.NewFromInt(x).Mod(decimal.NewFromInt(y)))
	//x, y := baseMethod.DecimalsToGrade(1123.7855)
	//fmt.Println(x, y)
	//	WebCralwer.Probability_ScoringRate("切沃", 3, true)
	//fmt.Println("总进球数:", WebCralwer.Calculate_sumGlobal("尤文图斯", 5, 1, "意甲"))
}

type EquationParam struct {
	homeTeam  string
	guestTeam string
}

//CreatGoldRateResult 创建进球率的结果
func CreatGoldRateResult() {
	var result_item Model.GoldRateResult
	//1.根据球队名称
	_ = result_item
}
