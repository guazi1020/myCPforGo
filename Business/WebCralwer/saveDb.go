package WebCralwer

import (
	"fmt"
	"myCPforGo/Com/DataBase/model"
	"myCPforGo/Com/DataBase/power"
	"myCPforGo/Model"
	"reflect"
	"strings"
	"time"

	tsgutils "github.com/typa01/go-utils"
)

const base_format = "2006-01-02 15:04:05"

const (
	userName = "root"
	password = "lijdlijd105"
	ip       = "49.235.158.254"
	port     = "3306"
	dbname   = "footballsp"
)

var comparam = model.DBParam{
	UserName: userName,
	Password: password,
	IP:       ip,
	Port:     port,
	Dbname:   dbname}
var enable model.Power

func init() {
	enable = power.ComMySQL{
		comparam}
}

//单一存储
func SaveOneGameInfo(game Model.Game) {

	var str_ string
	var str_value string
	t := reflect.TypeOf(game)
	v := reflect.ValueOf(game)
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			str_ = str_ + t.Field(i).Name + ","
			str_value = str_value + v.Field(i).String() + ","
			// fmt.Print(t.Field(i).Name)
			// fmt.Println(v.Field(i).String())
		}
	}
	str_ = strings.TrimRight(str_, ",")
	str_value = strings.TrimRight(str_value, ",")
	str_ = "insert game (" + str_ + ") values (" + str_value + ")"
	fmt.Println(str_)
}
func SaveDBTodey() {
	//fmt.Println(time.Now(time.Now().Year(), time.Now().Month(), time.Now().Day()))

	// t := time.Now()
	// str_time := t.Format(base_format)
	// fmt.Println(str_time)
	// //uuid的产生
	// fmt.Println(tsgutils.UUID())
	// fmt.Println(enable)
	//MysqlDemo_Select()
	var game Model.Game
	game.GIsfinish = "yes"
	SaveOneGameInfo(game)
}
func MysqlDemo_Insert() {
	str_sql := "insert game (UUID,Gnumber) values (?,?)"
	t := time.Now().Format(base_format)
	fmt.Println(enable.Exec(str_sql, tsgutils.UUID(), t))
}
func MysqlDemo_Select() {

	results := enable.Query("select * from game")
	//fmt.Println(results)
	for _, v := range results {
		for _k, _v := range v {
			fmt.Print(_k)
			fmt.Println(_v)
		}
		//fmt.Println(k)
		//fmt.Println(v)
	}
}
