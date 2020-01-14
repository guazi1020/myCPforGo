package main

import (
	"fmt"
	"myCPforGo/Business/WebCralwer"
)

func main() {
	fmt.Println("main, i am home ")
	/**/
	params := make(map[string]string)
	params["code"] = "201"
	params["ajax"] = "true"
	//WebCralwer.SaveWebByDate("2020-01-01", "", params)
	//fmt.Println(WebCralwer.Calculate_ScoringRate("切沃", 6))
	//fmt.Println(baseMethod.Compoundrate(152756, 0.0385, 24))
	//fmt.Println(WebCralwer.Calculate_AveGlobal("切沃", 3, true))
	//fmt.Println(baseMethod.CountFactorial(1))
	fmt.Println(WebCralwer.Probability_ScoringRate("切沃", 3))
}
