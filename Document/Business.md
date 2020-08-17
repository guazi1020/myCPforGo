目录说明

# 	1 Business

​	业务逻辑均在此目录下

##      	1.1 CPHttp

http接口的注册以及业务实现方法

##     	 1.2 WebCralwer

数据爬虫业务逻辑实现

# 	2 Com

## 		2.1 baseMethod

基本的方法实现

### 			2.2 comAddress.go

- RemoteIp(req *http.Request) string

  - req：request参数

  *获取当前访问者的IP地址*

- Ip2long(ipstr string) uint32

  - ipstr：需要被转换的信息
  
  *Ip2long 将 IPv4 字符串形式转为 uint32*
  
- LocalIPv4s() ([]string, error)

   *获取本机内网地址*

- GetNetIP() string

  *获取本机外网地址*

  ### 	2.3 comBandom.go

- (bandom *Bandom) CreatRandomInt() int  

  *生成Int随机数*

- (bandom *Bandom) CreatRandomFloat(m int) float64 

  *生成float64随机数(m:数据保留数量)*

- ChangeNumber(f float64, m int) string

  *保留小数点后几位方法(f:目标数据,m:数据保留位数)*

  ### 	2.4 comBasic.go

- CalculateGameResult(str string) string
  *计算比分结果 CalculateGameResult("1-0") 3*
  
  ###  2.5 comChannelsLimitCount.go
- (l *LimitRate) LimitChannel() bool 
- (l *LimitRate) SetRate(r int)
- (l *LimitRate) GetRate() int 
###     2.6 comGeneral.go
- MyPow(x float64, n int) float64
  *幂次方式*
  - x底数 
  - n幂次数
- DecimalsToGrade(source float64) (int64, int64)
  *小数变分数*
  - source 需要变换的小数值
  - return 第一个为分子；第二个为分母
- Compoundrate(source float64, rate float64, number int) float64
  *Compoundrate 计算复利*
  - source 本金
  - rate 利率
  - number 年份
  - return结果
- CountMultiplyingsqrt(source float64, number float64) float64
  *幂次方计算*
- CountMultiplying(source float64, number int) float64
  *幂次方计算*
- CountFactorial(source int) int 
  *阶乘*
  - 源数据


## 	3 commerr	
  ### 3.1 check.go
  - CheckErr(err error)
    CheckErr 容错
  ### 3.2 confOperate.go


## 	4 Database
### 4.1 /model/database.go
  接口文件定义数据库操作方法
### 4.2 /power/comMySQL.go
  MySQL数据库实现数据库操作
  -  (comMySQL ComMySQL) Query(a string, args ...interface{}) map[int]map[string]string 
    Query mysql实现查询语句
    - a SQL语句
    - args params
    - 返回为结果数据
 - (comMySQL ComMySQL) Exec(str_sql string, args ...interface{}) int64
    执行SQL语句
    -  comMySQL 执行语句
    - 返回结果数据 
# 	5 Config
  config.json 
  配置文件

# 	6 Document

# 	7 Interface

# 	8 Model

# 	9 Test