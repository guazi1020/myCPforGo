package Model

type GameAllBasic struct {
	//主键 UUID
	UUID string
	//赛事
	GAleague string
	//轮数
	GARound string
	//比赛时间
	GADate string
	//比赛状态
	GAIsFinished string
	//主队排名
	GAHomeRank string
	//主队名称
	GAHomeName string
	//客队排名
	GAGuestRank string
	//客队名称
	GAGuestName string
	//比分
	GAresultScore string
	//半场比分
	GAresultHalf string
	//比赛结果
	GAresult string

	//GspWin sp胜
	GAspWin string
	//sp平
	GAspTie string
	//sp负
	GAspDefeat string
	//计算的E值
	GAE string
	/*
		非业务性字段
	*/
	//创建时间
	CreateDate string
	//创建的IP
	CreateIP string
}
