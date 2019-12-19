package WebCralwer

import (
	"fmt"
	"time"
)

const base_format = "2006-01-02 15:04:05"

func SaveDBTodey() {
	//fmt.Println(time.Now(time.Now().Year(), time.Now().Month(), time.Now().Day()))
	t := time.Now()
	str_time := t.Format(base_format)
	fmt.Println(str_time)
	//uuid的产生
	
}
