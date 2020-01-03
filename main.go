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
	WebCralwer.SaveWebByDate("2020-01-03", "2020-01-03", params)
	//WebCralwer.SaveWebByDate("", "", params)
}
