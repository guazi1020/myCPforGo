package ImpMethod

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Model"
	"os"
)

type TTest struct {
	name string
}

type GetGameDataOne struct {
	StrYear string
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear() []Model.GameAllBasic {

	// filePtr, err := os.Open("../Config/config.json")
	// if err != nil {
	// 	fmt.Println("Open file failed [Err:%s]", err.Error())
	// }
	// decoder := json.NewDecoder(filePtr)
	// fmt.Println(decoder)
	// defer filePtr.Close()

	path, _ := os.Getwd()
	path += "\\Config\\configtest.json"

	filePtr, err := os.Open(path)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	decoder := json.NewDecoder(filePtr)
	var ttest []TTest
	err = decoder.Decode(&ttest)
	fmt.Println(ttest)
	defer filePtr.Close()
	var games []Model.GameAllBasic

	return games
}
