package main

import (
	"fmt"
	"myCPforGo/Com/baseMethod"
)

func main() {
	fmt.Println("main, i am home ")
	/**/
	params := make(map[string]string)
	params["code"] = "201"
	params["ajax"] = "true"
	//WebCralwer.SaveWebByDate("2018-10-02", "", params)
	//fmt.Println(WebCralwer.Calculate_ScoringRate("切沃", 2))
	fmt.Println(baseMethod.Compoundrate(152756, 0.0385, 24))

}
