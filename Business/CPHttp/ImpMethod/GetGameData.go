package ImpMethod

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Model"
	"os"
)

type TTest struct {
	Name string
}

type GetGameDataOne struct {
	Path string
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear() []Model.GameAllBasic {

	// path, _ := os.Getwd()
	// path += "\\Config\\configtest.json"

	filePtr, err := os.Open(getGameData.Path)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	decoder := json.NewDecoder(filePtr)
	var ttest TTest
	err = decoder.Decode(&ttest)
	fmt.Println(ttest.Name)
	defer filePtr.Close()
	var games []Model.GameAllBasic

	return games
}
