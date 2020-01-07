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
	WebCralwer.SaveWebByDate("2018-04-03", "", params)
	//WebCralwer.SaveWebByDate("", "", params)
	//fmt.Println(WebCralwer.IsOnly("2020-01-03"))
	//WebCralwer.MysqlDemo_Select()
}
