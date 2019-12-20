package main

import (
	"fmt"
	"myCPforGo/Business/WebCralwer"
)

func main() {
	fmt.Println("main, i am home ")
	WebCralwer.GetWeb()
	//WebCralwer.SaveDBTodey()
}
