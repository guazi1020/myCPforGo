package main

import (
	"fmt"
	"html"
	"log"
	"myCPforGo/Business/WebCralwer"
	"net/http"
)

func main() {
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
	var ishome int     //主客场
	var _ishome string //翻译临时主客场
	var league []string
	team = "多特蒙德"
	goals = 0
	num = 20
	ishome = 2

	switch ishome {
	case 0:
		_ishome = "主客场"
	case 1:
		_ishome = "主场"
	case 2:
		_ishome = "客场"
	}
	league = append(league, "德甲")
	//	league = append(league, "欧洲杯")

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
