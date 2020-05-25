# myCPforGo V0.01
## 目标
通过爬虫爬数据到数据库，并通过数学模型进行进行预测。
## 实现思路
找到数据网站，将数据爬下来，通过goquery对象化存入数据库中。
## API（示例）
#### 示例1
    使用方法：根据时间去爬网站数据，存到数据库中
    示例1，爬取2020-05-20到现在的网站数据
    '''
     params := make(map[string]string)
	 params["code"] = "201"
	 params["ajax"] = "true"
	 WebCralwer.SaveWebByDate("2020-05-20", "", params)
    '''
