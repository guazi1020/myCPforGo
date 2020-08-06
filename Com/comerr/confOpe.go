package comerr

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	strConfigFilePath = "./Config/configtest.json"
)

type Config struct {
	name string
}

func ReadConfig() {
	var config []Config
	config_file, err := os.Open(strConfigFilePath)
	if err != nil {
		panic("Failed to")
	}
	decoder := json.NewDecoder(config_file)
	err = decoder.Decode(&config)
	defer config_file.Close()
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		fmt.Println("Decoder success")
		fmt.Println(config)
	}
}
