package Model

type GameNow struct {
	GameInfo              Game
	GameE                 float64
	HomeScoringRate0      float64
	HomeScoringRate1      float64
	HomeScoringRate2      float64
	HomeScoringRate3      float64
	HomeScoringRateOther  float64
	GuestScoringRate0     float64
	GuestScoringRate1     float64
	GuestScoringRate2     float64
	GuestScoringRate3     float64
	GuestScoringRateOther float64

	NumberOfRound int
	LeagueName    string

	Gamestatistics GameStatistics
}
