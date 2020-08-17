package comerr

import (
	"fmt"
)

const (
	strConfigFilePath = "./Config/configtest.json"
)

type Configs struct {
	name string
}

func ReadConfig() {
	fmt.Println("ReadConfig")
	// var configs []Configs
	// config_file, err := os.Open(strConfigFilePath)
	// if err != nil {
	// 	panic("Failed to")
	// }
	// decoder := json.NewDecoder(config_file)
	// err = decoder.Decode(&configs)
	// defer config_file.Close()
	// if err != nil {
	// 	fmt.Println("Decoder failed", err.Error())

	// } else {
	// 	fmt.Println("Decoder success")
	// 	fmt.Println(configs)
	// }
}
