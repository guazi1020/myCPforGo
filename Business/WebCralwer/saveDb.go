package WebCralwer

import (
	"context"
	"fmt"
	"myCPforGo/Com/DataBase/model"
	"myCPforGo/Com/DataBase/power"
	"myCPforGo/Com/baseMethod"
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
func SaveOneGameInfo(game Model.Game, ctx context.Context, key int) {
	//这个功能只执行3.5s
	_, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(3500))
	defer func() {
		cancel()
		//close(ch)
	}()
	var str_ string
	//var str_value string
	var str_place string
	if (game == Model.Game{}) {
		// 如果对象是空的
		return
	} else {
		// 如果对象不为空,working
		t := reflect.TypeOf(game)

		game.CreateDate = time.Now().Format(base_format)
		game.CreateIP = baseMethod.GetNetIP()
		game.UUID = tsgutils.UUID()
		v := reflect.ValueOf(game)
		var pInterface []interface{} = make([]interface{}, v.NumField())
		for i := 0; i < t.NumField(); i++ {
			str_place += "?,"
			pInterface[i] = v.Field(i).String()
			str_ = str_ + t.Field(i).Name + ","
		}
		str_ = strings.TrimRight(str_, ",")
		str_place = strings.TrimRight(str_place, ",")

		str_sql := "insert into game (" + str_ + ") values (" + str_place + ")"
		_ = str_sql

		str_c := "insert into game values(" + str_place + ")"
		for v, k := range pInterface {
			fmt.Println(v, k)
		}
		// var pp []interface{} = make([]interface{}, 22)
		// pp[0] = tsgutils.UUID()
		// pp[1] = "周四001"
		if enable.Exec(str_c, pInterface...) == 1 {
			fmt.Println(key)
		}

		// _ = pInterface
		// str_ := "insert into game ( UUID ) values (?)"
		// var pp []interface{} = make([]interface{}, 1)
		// pp[0] = tsgutils.UUID()
		// //	fmt.Println(str_, pp)
		// if enable.Exec(str_, pp...) == 1 {
		// 	fmt.Println(key)
		// }
	}
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
