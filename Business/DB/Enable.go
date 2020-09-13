package DB

import (
	"myCPforGo/Com/DataBase/model"
	"myCPforGo/Com/DataBase/power"
	"myCPforGo/Config"
)

func Case() model.Power {
	configs := Config.ReadConfig()
	var comparam = model.DBParam{
		UserName: configs.Content[0].Book.UserName,
		Password: configs.Content[0].Book.Root,
		IP:       configs.Content[0].Book.Ip,
		Port:     configs.Content[0].Book.Port,
		Dbname:   configs.Content[0].Book.Dbname}
	return power.ComMySQL{
		comparam}
}
