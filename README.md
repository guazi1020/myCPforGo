# myCPforGo V0.01
## 目标
通过爬虫爬数据到数据库，并通过数学模型进行进行预测。
## 实现思路
找到数据网站，将数据爬下来，通过goquery对象化存入数据库中。
## API（示例）
### 1.根据时间去爬网站数据，存到数据库中

**示例**
爬取2020-05-20到现在的网站数据
   ```
   params := make(map[string]string)
	 params["code"] = "201"
	 params["ajax"] = "true"
	 WebCralwer.SaveWebByDate("2020-05-20", "", params)
  ```
### 2.预测进球率-预测某一支球队的进球可能性

预测的逻辑为：	
公式：
>P（X）=（M^X/X!)*e^(-M)；P (0) = e^(-M)

*M为球队场均进球数
X为期望进球值
e为常实数2.718
p(x)为最终进球概率*

**param**

team:球队,exceptGlobals 预测进球数,lastNumber 最近几场,isHome 是否是主场(0,全部;1,主场；2,客场),league 赛制名称

**方法名**

func Probability_ScoringRate(team string, exceptGlobals int, lastNumber int, isHome int, league ...string) float64 {}

**示例**
```
var team string    //球队名称
	var goals int      //进球数
	var num int        //几场比赛
	var ishome int     //主客场d
	var _ishome string //翻译临时主客场
	var league []string
	team = "多特蒙德"
	goals = 1
	num = 20
	ishome = 1

	switch ishome {
	case 0:
		_ishome = "主客场"
	case 1:
		_ishome = "主场"
	case 2:
		_ishome = "客场"
	}
	league = append(league, "德甲")
fmt.Printf("%s,进%d个球，范围为最近%d场%s，赛制为%s的情况下的可能性为：%f\n", team, goals, num, _ishome, league, WebCralwer.Probability_ScoringRate(team, goals, num, ishome, league...))
```
