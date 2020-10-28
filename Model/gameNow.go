package Model

type GameNow struct {
	//赛事基本信息
	GameInfo Game
	//GameE E值计算
	GameE float64
	//HomeScoringRate0 主队进零个概率
	HomeScoringRate0 float64
	//HomeScoringRate1 主队进1个的概率
	HomeScoringRate1 float64
	//HomeScoringRate2 主队进2个的概率
	HomeScoringRate2 float64
	//HomeScoringRate3 主队进3个的概率
	HomeScoringRate3 float64
	//HomeScoringRateOther 主队进多个的概率
	HomeScoringRateOther float64
	//GuestScoringRate0 客队进零个的概率
	GuestScoringRate0 float64
	//GuestScoringRate1 客队进1个的概率
	GuestScoringRate1 float64
	//GuestScoringRate2 客队进2个的概率
	GuestScoringRate2 float64
	//GuestScoringRate3 客队进3个的概率
	GuestScoringRate3 float64
	//GuestScoringRateOther 客队进多个的概率
	GuestScoringRateOther float64
	//比赛的轮数
	NumberOfRound int
	//赛事名称
	LeagueName string
	//赛事统计
	Gamestatistics GameStatistics
}
