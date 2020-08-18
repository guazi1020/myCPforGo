package Config

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Com/comerr"
	"os"
)

type ModelConfig struct {
	Title   string
	Vision  string
	Content []Dataset
}
type Book struct {
	UserName string
	Root     string
	Ip       string
	Port     string
	Dbname   string
}
type Dataset struct {
	Book     Book
	Descript string
	Type     string
	Vison    string
	Endtime  string
}

//ReadConfig 读取配置文件
func ReadConfig() {
	path, _ := os.Getwd()
	path += "\\Config\\config.json"

	filePtr, err := os.Open(path)
	if err != nil {
		comerr.CheckErr(err)
	}
	decoder := json.NewDecoder(filePtr)
	var configs ModelConfig
	err = decoder.Decode(&configs)
	fmt.Println(configs.Content[0].Book.Dbname)
	defer filePtr.Close()

}
