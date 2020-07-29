/**
** 数据库操作的前一层
**/

package WebCralwer

import (
	"context"
	"fmt"
	"log"
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

// const (
// 	userName = "root"
// 	password = "LijdLijd*105"
// 	//password = "lijdlijd105"
// 	ip = "49.235.158.254"
// 	//port     = "3306"
// 	port   = "38160"
// 	dbname = "footballsp"
// )

const (
	userName = "root"
	password = "LijdLijd*105"
	//password = "lijdlijd105"
	ip = "cdb-6zkfgy0x.gz.tencentcdb.com"
	//port     = "3306"
	port   = "10155"
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

	var results map[int]map[string]string
	if len(league) == 0 {

		switch ishome {
		default:
			break
		case 0:
			//results = enable.Query("select * from game WHERE GhomeName=? or GguestName=? order by Gyear desc LIMIT ?", team, team, count)
			results = enable.Query("select * from GameAllBasic WHERE GAhomeName=? or GguestName=? order by GADate desc LIMIT ?", team, team, count)
		case 1:
			//results = enable.Query("select * from game WHERE GhomeName=?  order by Gyear desc LIMIT ?", team, count)
			results = enable.Query("select * from GameAllBasic WHERE GAhomeName=?  order by GADate desc LIMIT ?", team, count)
		case 2:
			//results = enable.Query("select * from game WHERE  GguestName=? order by Gyear desc LIMIT ?", team, count)
			results = enable.Query("select * from GameAllBasic WHERE  GAguestName=? order by GADate desc LIMIT ?", team, count)
		}

	}
	if len(league) > 0 {
		var leagues string
		for _, item := range league {
			leagues += "'" + item + "',"
		}
		leagues = strings.TrimRight(leagues, ",")
		//fmt.Println("leagues:", leagues)
		var sql string
		switch ishome {
		default:
			break
		case 0:
			//sql = fmt.Sprintf("select * from game WHERE GhomeName=? or GguestName=? and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			sql = fmt.Sprintf("select * from GameAllBasic WHERE GAhomeName=? or GAguestName=? and Gleague in (%s) order by GADate desc LIMIT ?", leagues)
			results = enable.Query(sql, team, team, count)
		case 1:
			//sql = fmt.Sprintf("select * from game WHERE GhomeName=?  and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			sql = fmt.Sprintf("select * from GameAllBasic WHERE GAhomeName=?  and GAleague in (%s) order by GADate desc LIMIT ?", leagues)
			results = enable.Query(sql, team, count)
		case 2:
			//sql = fmt.Sprintf("select * from game WHERE  GguestName=? and Gleague in (%s) order by Gyear desc LIMIT ?", leagues)
			sql = fmt.Sprintf("select * from GameAllBasic WHERE  GAguestName=? and GAleague in (%s) order by GADate desc LIMIT ?", leagues)
			results = enable.Query(sql, team, count)
		}
		//fmt.Println(sql)
	}

	return results
}

//SaveOneGameInfo 单一存储
//game:模型,ctx:上下文规定deadline,key:协程编号
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

func SaveOneGameAllBaiscInfo(gameAll Model.GameAllBasic, ctx context.Context, key int) {
	//这个功能只执行1.5s
	_, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(1500))
	defer func() {
		cancel()
	}()

	if (gameAll == Model.GameAllBasic{}) {
		// 如果对象是空的
		return
	} else {

		/*
			构建str_sql
		*/
		var str_place string //str_place 标识符
		gameAll.CreateDate = time.Now().Format(base_format)
		gameAll.CreateIP = baseMethod.GetNetIP()
		gameAll.UUID = tsgutils.UUID()

		t := reflect.TypeOf(gameAll)
		v := reflect.ValueOf(gameAll)
		var pInterface []interface{} = make([]interface{}, v.NumField())
		//构建pInterface参数[]interface{}
		for i := 0; i < t.NumField(); i++ {
			str_place += "?,"
			pInterface[i] = v.Field(i).String()

		}

		str_place = strings.TrimRight(str_place, ",")
		str_c := "insert into GameAllBasic values(" + str_place + ")" //str_c 格式:insert into game values(?,?,?...)
		//fmt.Println(str_c)
		//存储
		if enable.Exec(str_c, pInterface...) == 1 {
			fmt.Println(key, "finished")
		}
	}
}

//IsOnly 判断是否有重复的值
//ture 有重复的 false 无重复的
func IsOnly(Gyear string, tableName string, primaryField string) bool {

	str_sql := "select count({primaryField}) num from {tableName} where {primaryField}=?"
	r := strings.NewReplacer("{tableName}", tableName, "{primaryField}", primaryField)
	str_sql = r.Replace(str_sql)
	i, _ := strconv.Atoi(enable.Query(str_sql, Gyear)[0]["num"])
	if i > 0 {
		return true //有重复的,true
	}
	return false //没有重复的, false
}

func SaveLeague(league Model.League) {
	if (league == Model.League{}) {
		return
	}
	league.UUID = tsgutils.UUID()
	str_place, pInterface := ModeltoString(league, "DicLeague")
	if enable.Exec(str_place, pInterface...) == 1 {
		fmt.Println("finished")
	}
}

/*
Savecomm 通用的保存数据
	Param |mmodel:实体类,tablename:表名|
*/
func SaveComm(mmodel interface{}, tablename string) {

	if tablename == "" {
		log.Println("表名没写")
		return
	}
	str_place, pInterface := ModeltoString(mmodel, tablename)

	if pInterface == nil {
		log.Println("model是空的")
		return
	}
	if enable.Exec(str_place, pInterface...) == 1 {
		log.Println("finished")
	}
}

//ModeltoString 模型转换为拆分占位符和param语句,为了给insert用
//model:模型 tableName:操作的数据库表对象
//返回1.占位符 2.param
/*Demo
str_place, pInterface := ModeltoString(league, "DicLeague")
enable.Exec(str_place,pInterface...)
*/
func ModeltoString(model interface{}, tableName string) (string, []interface{}) {
	var str_place string //字符串
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	var pInterface []interface{} = make([]interface{}, v.NumField())
	//构建pInterface参数[]interface{}
	for i := 0; i < t.NumField(); i++ {
		str_place += "?,"
		pInterface[i] = v.Field(i).String()
	}
	str_place = strings.TrimRight(str_place, ",")
	str_insert := "insert into " + tableName + " values(" + str_place + ")"
	return str_insert, pInterface
}

//MakeGameStatistics 历史E+-0.05上下同类league的统计,填充game模型
func MakeGameStatistics(game Model.GameNow) Model.GameNow {

	//conRank Rank条件
	var conRank string
	homeranki, _ := strconv.Atoi(game.GameInfo.GhomeRank)
	guestranki, _ := strconv.Atoi(game.GameInfo.GguestRank)
	if homeranki < guestranki {
		conRank = " GAhomeRank < GAguestRank "
	} else {
		conRank = " GAhomeRank >= GAguestRank "
	}
	if len(conRank) == 0 {
		conRank = " 1=1 "
	}

	//conSP conSP条件
	var conSP string
	homeSPi, _ := strconv.Atoi(game.GameInfo.GspWin)
	guestSPi, _ := strconv.Atoi(game.GameInfo.GspDefeat)
	if homeSPi < guestSPi {
		conSP = " GAspWin < GAspDefeat "
	} else {
		conSP = " GAspWin >= GAspDefeat "
	}
	if len(conSP) == 0 {
		conSP = " 1=1 "
	}

	str := "SELECT COUNT(*) as 'GCount', sum(CASE GAresult WHEN '3' THEN 1 ELSE 0 END) as 'GWinNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '3' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GWinDC', sum(CASE GAresult WHEN '1' THEN 1 ELSE 0 END) as 'GTieNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '1' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GTietDC', sum(CASE GAresult WHEN '0' THEN 1 ELSE 0 END) as 'GDefeatNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '0' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GDefeatDC' FROM GameAllBasic where GAleague=? and GAE>? and GAE<? and GADate>? "
	str = str + " and " + conRank + " and " + conSP
	var params []interface{}
	params = append(params, game.GameInfo.Gleague)
	//beginE, _ := strconv.ParseFloat(game.GameE, 64)
	params = append(params, baseMethod.ChangeNumber(game.GameE-0.05, 3))
	//params = append(params, 0)
	//strconv.FormatFloat(v, 'E', -1, 64)/
	params = append(params, baseMethod.ChangeNumber(game.GameE+0.05, 3))
	//params = append(params, 4)
	//fmt.Println(baseMethod.ChangeNumber(game.GameE+0.01, 2))
	params = append(params, "2019")

	results := enable.Query(str, params...)
	if len(results) > 0 {
		game.Gamestatistics.GCount = results[0]["GCount"]
		game.Gamestatistics.GDefeatDC = results[0]["GDefeatDC"]
		game.Gamestatistics.GDefeatNumber = results[0]["GDefeatNumber"]
		game.Gamestatistics.GTieNumber = results[0]["GTieNumber"]
		game.Gamestatistics.GTietDC = results[0]["GTietDC"]
		game.Gamestatistics.GWinDC = results[0]["GWinDC"]
		game.Gamestatistics.GWinNumber = results[0]["GWinNumber"]
	}
	return game
}

//clearRepeatInfo 清除多余的数据
func ClearRepeatInfo() {
	//1.将多余的数据存储到toDel表中
	str := "INSERT INTO toDel ( SELECT MIN(UUID),	GAleague,	GARound,	GAHomeName,	GAHomeRank,	GAGuestName,	GAGuestRank,	GAresultScore FROM	GameAllBasic WHERE	GAHomeRank != '' GROUP BY	GARound,	GAHomeName,	GAHomeRank,	GAGuestName,	GAGuestRank,	GAresultScore,	GAleague HAVING	count(*) > 1)"
	//fmt.Println(enable.Exec(str))
	if enable.Exec(str) > 0 {
		//2.删除gamebasic的数据。
		strDel := "DELETE FROM GameAllBasic WHERE	uuid IN (SELECT uuid FROM toDel)"
		if enable.Exec(strDel) > 0 {
			strDeltoDel := "delete from toDel"
			if enable.Exec(strDeltoDel) > 0 {
				fmt.Println("已删除重复的数据")
			}
		}
	}

}

/*SearchCom 标准查找方法
 */
func SearchCom(str_sql string) map[int]map[string]string {
	return enable.Query(str_sql)
}
func MysqlDemo_Insert() {
	str_sql := "insert game (UUID,Gnumber) values (?,?)"
	t := time.Now().Format(base_format)
	fmt.Println(enable.Exec(str_sql, tsgutils.UUID(), t))
}
func MysqlDemo_Select() {
	str := "SELECT COUNT(*) as 'GCount', sum(CASE GAresult WHEN '3' THEN 1 ELSE 0 END) as 'GWinNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '3' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GWinDC', sum(CASE GAresult WHEN '1' THEN 1 ELSE 0 END) as 'GTieNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '1' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GTietDC', sum(CASE GAresult WHEN '0' THEN 1 ELSE 0 END) as 'GDefeatNumber', CONCAT(CAST(round((sum(CASE GAresult WHEN '0' THEN 1 ELSE 0 END)/COUNT(*))*100,3) AS CHAR),'%') AS 'GDefeatDC' FROM GameAllBasic where GAleague='英冠' and GAE>'1.22' and GAE<'1.23' and GADate>'' order BY GADate DESC "
	results := enable.Query(str)
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
