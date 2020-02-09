package WebCralwer

import (
	"context"
	"fmt"
	"myCPforGo/Com/DataBase/model"
	"myCPforGo/Com/DataBase/power"
	"myCPforGo/Com/baseMethod"
	"myCPforGo/Model"
	"reflect"
	"strconv"
	"strings"
	"time"

	tsgutils "github.com/typa01/go-utils"
)

const base_format = "2006-01-02 15:04:05"

const (
	userName = "root"
	password = "LijdLijd*105"
	//password = "lijdlijd105"
	ip = "49.235.158.254"
	//port     = "3306"
	port   = "38160"
	dbname = "footballsp"
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

//SearchForGame 根据game查找内容
//team 队名 count 前n条,ishome 0,全部 1,主场 2,客场
func SearchForGame(team string, count int, ishome int, league ...string) map[int]map[string]string {
	fmt.Println(league)
	var results map[int]map[string]string
	if len(league) == 0 {

		switch ishome {
		default:
			break
		case 0:
			results = enable.Query("select * from game WHERE GhomeName=? or GguestName=? order by Gyear desc LIMIT ?", team, team, count)
		case 1:
			results = enable.Query("select * from game WHERE GhomeName=?  order by Gyear desc LIMIT ?", team, count)
		case 2:
			results = enable.Query("select * from game WHERE  GguestName=? order by Gyear desc LIMIT ?", team, count)
		}

	}
	if len(league) > 0 {
		var leagues string
		for _, item := range league {
			leagues += "'" + item + "',"
		}
		leagues = strings.TrimRight(leagues, ",")
		fmt.Println("leagues:", leagues)
		var sql string
		switch ishome {
		default:
			break
		case 0:
			sql = fmt.Sprintf("select * from game WHERE GhomeName=? or GguestName=? and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			results = enable.Query(sql, team, team, count)
		case 1:
			sql = fmt.Sprintf("select * from game WHERE GhomeName=?  and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			results = enable.Query(sql, team, count)
		case 2:
			sql = fmt.Sprintf("select * from game WHERE  GguestName=? and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			results = enable.Query(sql, team, count)
		}
		fmt.Println(sql)
	}

	return results
}

//SaveOneGameInfo 单一存储
func SaveOneGameInfo(game Model.Game, ctx context.Context, key int) {
	//这个功能只执行1.5s
	_, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(1500))
	defer func() {
		cancel()
	}()

	if (game == Model.Game{}) {
		// 如果对象是空的
		return
	} else {

		/*
			构建str_sql
		*/
		var str_place string //str_place 标识符
		game.CreateDate = time.Now().Format(base_format)
		game.CreateIP = baseMethod.GetNetIP()
		game.UUID = tsgutils.UUID()

		t := reflect.TypeOf(game)
		v := reflect.ValueOf(game)
		var pInterface []interface{} = make([]interface{}, v.NumField())
		//构建pInterface参数[]interface{}
		for i := 0; i < t.NumField(); i++ {
			str_place += "?,"
			pInterface[i] = v.Field(i).String()
		}

		str_place = strings.TrimRight(str_place, ",")
		str_c := "insert into game values(" + str_place + ")" //str_c 格式:insert into game values(?,?,?...)
		//fmt.Println(str_c)
		//存储
		if enable.Exec(str_c, pInterface...) == 1 {
			fmt.Println(key, "finished")
		}
	}
}

//IsOnly 判断是否有重复的值
//ture 有重复的 false 无重复的
func IsOnly(Gyear string) bool {
	str_sql := "select count(Gyear) num from game where Gyear=?"
	i, _ := strconv.Atoi(enable.Query(str_sql, Gyear)[0]["num"])
	if i > 0 {
		return true //有重复的,true
	}
	return false //没有重复的, false
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
