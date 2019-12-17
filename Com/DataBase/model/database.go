package model

//DBParam 数据库参数
type DBParam struct {
	UserName string
	Password string
	IP       string
	Port     string
	Dbname   string
}

//Power 能力接口
type Power interface {

	/*
		Query 查询
		return 结果集
	*/
	Query(string, ...interface{}) map[int]map[string]string
	/*
		Exec 非查询SQL执行
		return 返回成功个数
		@string sql语句
		@...interface{} 动态param
	*/
	Exec(string, ...interface{}) int64
}
