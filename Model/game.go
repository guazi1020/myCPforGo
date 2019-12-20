package Model

import "fmt"

type Game struct {
	//赛事编号
	Gnumber string
	//日期
	Gdata string
	//时间
	Gtime string
	//赛制
	Gleague string
	//轮
	Gleaguenumber string
	//是否已经完赛
	GIsfinish string
	//主队排名
	GhomeRank string
	//主队名称
	GhomeName string
	//客队排名
	GguestRank string
	//客队名称
	GguestName string
	//结果
	Gresult string
	//GspWin sp胜
	GspWin string
	//sp平
	GspTie string
	//sp负
	GspDefeat string
	//全场比分
	Gscore string
	//半场比分
	GhalfSource string
	//红牌数量
	GredQuantities string
	//让球数
	GletCount string
	//比赛结果
	GresultScore string
	//客队红牌
	GredQuantitlesGuest string
	//半场比分
	GresultHalfScore string
}

// SaveGametoDB 把比赛数据放到数据库中
func (game *Game) SaveGametoDB() {
	fmt.Println(game.Gnumber)
}