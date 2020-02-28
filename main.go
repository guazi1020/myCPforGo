package main

import (
	"encoding/json"
	"fmt"
	"html"
	"myCPforGo/Business/WebCralwer"
	"myCPforGo/Model"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	StartHttp()
	//Equation()
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
}

//StartHttp 开始启动httpweb
func StartHttp() {
	router := mux.NewRouter().StrictSlash(true)
	//注册
	router.HandleFunc("/", HandleIndex)
	router.HandleFunc("/app", HandleDemoIndex)
	router.HandleFunc("/app/{id}", HandleDemoShow)

	fmt.Println("Main task")
	http.ListenAndServe(":8080", router)
	//log.Fatal(http.ListenAndServe(":8080", router))
}

//HandleIndex /index
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
}

//HandleDemoIndex /app
func HandleDemoIndex(w http.ResponseWriter, r *http.Request) {
	domains := Games{
		Model.Game{UUID: "a"},
		Model.Game{UUID: "b"},
	}
	_ = domains
	json.NewEncoder(w).Encode(WebCralwer.SearchForGame("AC米兰", 10, 0))
}

// fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
// fmt.Fprintf(w, "this is app")
type Games []Model.Game

//HandDemoShow /app/{i}
func HandleDemoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Domain Show:%q", id)
}
