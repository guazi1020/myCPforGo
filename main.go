package main

import (
	"fmt"
	"myCPforGo/Business/WebCralwer"
)

func main() {
	fmt.Println("main, i am home ")
	/*
		params := make(map[string]string)
		params["code"] = "201"
		params["date"] = "2019-11-23"
		params["ajax"] = "true"
		WebCralwer.SaveWeb(params)
	*/

	WebCralwer.SaveWebByDate("2019-11-23", "2019-12-26")
}
