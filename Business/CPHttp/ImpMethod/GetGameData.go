package ImpMethod

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Model"
	"os"
)

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
	path += "\\Config\\config.json"
	fmt.Println(path)
	filePtr, err := os.Open(path)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	decoder := json.NewDecoder(filePtr)
	fmt.Println(decoder)
	defer filePtr.Close()
	var games []Model.GameAllBasic

	return games
}
