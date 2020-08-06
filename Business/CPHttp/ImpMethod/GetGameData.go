package ImpMethod

import (
	"fmt"
	"myCPforGo/Model"
	"os"
)

type GetGameDataOne struct {
}

//GetGameDataForYear 实现
func (getGameData GetGameDataOne) GetGameDataForYear(strYear string) []Model.GameAllBasic {

	// filePtr, err := os.Open("../Config/config.json")
	// if err != nil {
	// 	fmt.Println("Open file failed [Err:%s]", err.Error())
	// }
	// decoder := json.NewDecoder(filePtr)
	// fmt.Println(decoder)
	// defer filePtr.Close()

	path, _ := os.Getwd()
	fmt.Println("Path:", path)
	var games []Model.GameAllBasic
	fmt.Println("one running")
	return games
}
