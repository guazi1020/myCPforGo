package power

import (
	"database/sql"
	"myCPforGo/Com/DataBase/model"
	"myCPforGo/Com/comerr"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//ComMySQL mysql实体实现
type ComMySQL struct {
	model.DBParam
}

//DB mysql的连接变量

//Query mysql实现查询语句
func (comMySQL ComMySQL) Query(a string, args ...interface{}) map[int]map[string]string {
	result := make(map[int]map[string]string)
	var DB *sql.DB
	/*数据库操作*/
	DB = comMySQL.Open()
	stmt, err := DB.Prepare(a)
	if err != nil {
		//fmt.Println("数据库连接失败")
		comerr.CheckErr(err)
	}
	rows, err := stmt.Query(args...)
	comerr.CheckErr(err)

	/*
	* 单行数据地址关联和转换
	 */
	cols, _ := rows.Columns()                    //表的列list
	values := make([][]byte, len(cols))          //最终信息的存储(一行信息),[]byte存储信息
	addr_proxy := make([]interface{}, len(cols)) //地址代理
	for k, _ := range values {                   //地址关联,到时候数据会直接存到vuales中。
		addr_proxy[k] = &values[k]
	}

	i := 0            //初始行数
	for rows.Next() { //扫描
		rows.Scan(addr_proxy...)       //将一行内容装载
		row := make(map[string]string) //创建行
		for k, v := range values {     //将内容装到row里面去
			row[cols[k]] = string(v)
		}
		result[i] = row
		i++
	}
	DB.Close()
	stmt.Close()
	return result
}

//Del mysql实现执行sql
func (comMySQL ComMySQL) Exec(str_sql string, args ...interface{}) int64 {
	var DB *sql.DB
loop:
	DB = comMySQL.Open()
	stmt, err := DB.Prepare(str_sql)
	if err != nil {
		time.Sleep(3 * 1000)
		goto loop
	}
	comerr.CheckErr(err)
	//st := "insert into game (UUID,Gnumber,Gdata,Gtime,Gleague,Gleaguenumber,GIsfinish,GhomeRank,GhomeName,GguestRank,GguestName,Gresult,GspWin,GspTie,GspDefeat,GredQuantities,GletCount,GresultScore,GredQuantitlesGuest,GresultHalfScore,CreateDate,CreateIP) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := stmt.Exec(args...)
	comerr.CheckErr(err)
	num, err := result.RowsAffected()
	comerr.CheckErr(err)
	defer func() {
		if num != 0 {
			stmt.Close()
			DB.Close()
		}
	}()
	return num
}

//Open 实现连接数据库
func (comMysql ComMySQL) Open() *sql.DB {
	//组装账套
	path := strings.Join([]string{comMysql.UserName, ":", comMysql.Password, "@tcp(", comMysql.IP, ":", comMysql.Port, ")/", comMysql.Dbname, "?charset=utf8"}, "")
	//fmt.Println(path)
	//连接
	db, err := sql.Open("mysql", path)
	comerr.CheckErr(err)
	return db

}
