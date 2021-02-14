package Test

import (
	"os"
	"testing"
)

func TestA(t *testing.T) {
	// var httpMode HTTP.IGetGameData
	// htt pMode = ImpMethod.GetGameDataOne{Year: "2020"}
	// httpMode.GetGameDataForYear()
	//
	// Mf := Config.ReadConfig()
	// fmt.Printf(Mf.Title)
	path, _ := os.Getwd()
	path += "/Config/config.json"
	//fmt.Println("path")
	//dirPath := filepath.Dir("config.json")
	t.Log(path)
	// fmt.Println("ccc" + dirPath)
}
