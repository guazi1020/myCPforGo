package Model

import (
	"fmt"
	"reflect"
)

type Game struct {
	//主键 UUID
	UUID string
	//赛事编号
	Gnumber string
	//年份
	Gyear string
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
	//比赛结果
	GresultScore string
	//结果
	Gresult string
	//GspWin sp胜
	GspWin string
	//sp平
	GspTie string
	//sp负
	GspDefeat string
	//半场比分
	GresultHalfScore string
	//红牌数量
	GredQuantities string
	//客队红牌
	GredQuantitlesGuest string
	//让球数
	GletCount string
	/*
		非业务性字段
	*/
	//创建时间
	CreateDate string
	//创建的IP
	CreateIP string
	//E值
	GE string
}

// SaveGametoDB 把比赛数据放到数据库中
func (game *Game) SaveGametoDB() {
	fmt.Println(game.Gnumber)
}
func (game *Game) IsEmpty() bool {
	return reflect.DeepEqual(game, Game{})
}

//GoldPredictionItem 进球预测模型
type GoldPredictionItem struct {
	//goalnumber 进球数
	goalnumber string
	//rate 进球率
	rate string
}

type GoldRateResult struct {
	//Model.Game 参赛球队
	Game Game
	//result 进球结果
	Result []GoldPredictionItem
}
