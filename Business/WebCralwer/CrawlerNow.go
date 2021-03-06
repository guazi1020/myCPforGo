package WebCralwer

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Com"
	"myCPforGo/Model"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/tealeg/xlsx"
)

//GetEbyDateLiveStreaming 根据当前时间测算比赛（live-stream）
func GetEbyDateLiveStreaming(round int, strDate string) []Model.Game {

	// c := colly.NewCollector()
	games := []Model.Game{}
	params := make(map[string]string)
	params["code"] = "201"
	params["ajax"] = "true"
	params["date"] = time.Now().Format("2006-01-02")
	//params["date"] = strDate
	url := CompositionURL("http://live.zgzcw.com/ls/AllData.action", params)
	games = GetWeb(url, params) //获取今天比赛
	gameNows := []Model.GameNow{}
	for _, game := range games {
		gameNow := Model.GameNow{}
		gameNow.GameInfo = game
		ghomerank, _ := strconv.ParseFloat(game.GhomeRank, 32)
		gguestRank, _ := strconv.ParseFloat(game.GguestRank, 32)
		gspwin, _ := strconv.ParseFloat(game.GspWin, 32)
		gameNow.GameE = Calculate_E(gguestRank-ghomerank, gspwin)
		gameNow = MakeGameStatistics(gameNow)
		fmt.Println(gameNow.GameInfo.Gnumber)
		fmt.Println(gameNow.GameE, gameNow.GameInfo.GspWin, gameNow.GameInfo.GspTie, gameNow.GameInfo.GspDefeat, gameNow.Gamestatistics)
		gameNows = append(gameNows, gameNow)

	}
	OutToExcel(gameNows)
	// fmt.Println(string(data))

	return games
}

//GetEByDate 根据当前时间来测算每场的E值
func GetEByDate(round int) []Model.GameNow {
	//1.筛选符合条件的比赛
	//a.sp值均大于2.0
	c := colly.NewCollector()
	games := []Model.GameNow{}
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
			//fmt.Println(s.Html())

			//创建game
			game := Model.Game{}
			gameNow := Model.GameNow{}
			//t+mn为Gnumber
			t, _ := s.Attr("t")
			mn, _ := s.Attr("mn")
			game.Gnumber = t + mn
			//Gleague
			m, _ := s.Attr("m")
			game.Gleague = m
			//GhomeName
			game.GhomeName = s.Find("td.wh-4").Find("a").Text()
			game.GhomeRank = Com.RemoveBlank((Com.RemoveCharacter(Com.RemoveCharacter(s.Find("td.wh-4").Find("em").Text(), "", "]"), "", "[")))
			//Gguest
			game.GguestName = s.Find("td.wh-6").Find("a").Text()
			game.GguestRank = Com.RemoveBlank((Com.RemoveCharacter(Com.RemoveCharacter(s.Find("td.wh-6").Find("em").Text(), "", "]"), "", "[")))

			//EachWithBreak func(f func(int, *Selection) bool) *Selection
			//Each func(f func(int, *Selection)) *Selection

			tag := true
			s.Find("td.wh-8").Find("div.tz-area").Eq(0).Find("a").Each(func(i int, s *goquery.Selection) {

				if s.Text() == "未开售" {
					return
				}
				//ffloat, _ := strconv.ParseFloat(s.Text(), 32)
				switch i {
				case 0:
					game.GspWin = s.Text()

				case 1:
					game.GspTie = s.Text()

				case 2:
					game.GspDefeat = s.Text()

				default:
					return
				}
			})
			gspwin, _ := strconv.ParseFloat(game.GspWin, 32)
			ghomerank, _ := strconv.ParseFloat(game.GhomeRank, 32)
			gguestRank, _ := strconv.ParseFloat(game.GguestRank, 32)

			gspDefeat, _ := strconv.ParseFloat(game.GspDefeat, 32)
			gspDeTie, _ := strconv.ParseFloat(game.GspTie, 32)
			//判定都大于2.0
			if gspDefeat > 2 && gspwin > 2 && gspDeTie > 2 {
				tag = true
			}

			//计算E
			gameNow.GameInfo = game
			gameNow.GameE = Calculate_E(gguestRank-ghomerank, gspwin)

			//派出league

			// if gameNow.GameE != 0 && tag == true && (gameNow.GameE >  || gameNow.GameE < 0.923) {
			if gameNow.GameE != 0 && tag == true {
				//计算最近的进球率
				tnum := round
				gameNow.HomeScoringRate0 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 0, tnum, 1, gameNow.GameInfo.Gleague)
				gameNow.HomeScoringRate1 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 1, tnum, 1, gameNow.GameInfo.Gleague)
				gameNow.HomeScoringRate2 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 2, tnum, 1, gameNow.GameInfo.Gleague)
				gameNow.HomeScoringRate3 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 3, tnum, 1, gameNow.GameInfo.Gleague)
				gameNow.HomeScoringRateOther = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 4, tnum, 1, gameNow.GameInfo.Gleague)
				gameNow.GuestScoringRate0 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 0, tnum, 2, gameNow.GameInfo.Gleague)
				gameNow.GuestScoringRate1 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 1, tnum, 2, gameNow.GameInfo.Gleague)
				gameNow.GuestScoringRate2 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 2, tnum, 2, gameNow.GameInfo.Gleague)
				gameNow.GuestScoringRate3 = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 3, tnum, 2, gameNow.GameInfo.Gleague)
				gameNow.GuestScoringRateOther = Probability_ScoringRate(gameNow.GameInfo.GhomeName, 4, tnum, 2, gameNow.GameInfo.Gleague)

				gameNow = MakeGameStatistics(gameNow)
				//fmt.Println(gameNow.LeagueName, gameNow.GameE)
				games = append(games, gameNow)
			}

		})
		data, _ := json.Marshal(games)
		fmt.Println(string(data))
		OutToExcel(games)
	})
	c.Visit("http://cp.zgzcw.com/lottery/jchtplayvsForJsp.action?lotteryId=47&type=jcmini&issue=" + time.Now().Format("2006-01-02"))
	return games
}

//OutToExcel 根据model.GameNow的模型进行生成Excel
func OutToExcel(gamesNow []Model.GameNow) {
	//file := xlsx.NewFile()
	file, _ := xlsx.OpenFile("file.xlsx")
	sheet, err := file.AddSheet(time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("今日比赛已经计算过了。")
	}
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	//  gamenow := Model.GameNow{}
	//  gamenow.GameInfo.Gnumber
	//生成表头
	cell := row.AddCell()
	cell.Value = "比赛编号"
	cell = row.AddCell()
	cell.Value = "赛事"
	cell = row.AddCell()
	cell.Value = "主队"
	cell = row.AddCell()
	cell.Value = "主队排名"
	cell = row.AddCell()
	cell.Value = "客队"
	cell = row.AddCell()
	cell.Value = "客队排名"
	cell = row.AddCell()
	cell.Value = "胜"
	cell = row.AddCell()
	cell.Value = "平"
	cell = row.AddCell()
	cell.Value = "负"
	cell = row.AddCell()
	cell.Value = "E"
	cell = row.AddCell()
	cell.Value = "H0"
	cell = row.AddCell()
	cell.Value = "H1"
	cell = row.AddCell()
	cell.Value = "H2"
	cell = row.AddCell()
	cell.Value = "H3"
	cell = row.AddCell()
	cell.Value = "H4"
	cell = row.AddCell()
	cell.Value = "G0"
	cell = row.AddCell()
	cell.Value = "G1"
	cell = row.AddCell()
	cell.Value = "G2"
	cell = row.AddCell()
	cell.Value = "G3"
	cell = row.AddCell()
	cell.Value = "G4"

	cell = row.AddCell()
	cell.Value = "符合总比赛数"
	cell = row.AddCell()
	cell.Value = "胜场"
	cell = row.AddCell()
	cell.Value = "胜场占比"
	cell = row.AddCell()
	cell.Value = "平场"
	cell = row.AddCell()
	cell.Value = "平场占比"
	cell = row.AddCell()
	cell.Value = "负场"
	cell = row.AddCell()
	cell.Value = "负场占比"

	for _, gamenow := range gamesNow {
		dtrow := sheet.AddRow()
		dtrow.SetHeightCM(1)
		//比赛编号
		dtcell := dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.Gnumber
		//赛事
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.Gleague
		//主队
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GhomeName
		//主队排名
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GhomeRank
		//客队
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GguestName
		//客队排名
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GguestRank
		//胜
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GspWin
		//平
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GspTie
		//负
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.GameInfo.GspDefeat
		//E
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GameE, 'f', 6, 64)
		//H0
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.HomeScoringRate0, 'f', 6, 64)
		//H1
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.HomeScoringRate1, 'f', 6, 64)
		//H2
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.HomeScoringRate2, 'f', 6, 64)
		//H3
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.HomeScoringRate3, 'f', 6, 64)
		//H4
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.HomeScoringRateOther, 'f', 6, 64)
		//G0
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GuestScoringRate0, 'f', 6, 64)
		//G1
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GuestScoringRate1, 'f', 6, 64)
		//G2
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GuestScoringRate2, 'f', 6, 64)
		//G3
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GuestScoringRate3, 'f', 6, 64)
		//G4
		dtcell = dtrow.AddCell()
		dtcell.Value = strconv.FormatFloat(gamenow.GuestScoringRateOther, 'f', 6, 64)

		//GCount        string //数量
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GCount
		//GWinNumber    string //胜场
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GWinNumber
		//GWinDC        string //占比
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GWinDC
		//GTieNumber    string //平场
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GTieNumber
		//GTietDC       string //平比
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GTietDC
		//GDefeatNumber string //客胜场
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GDefeatNumber
		//GDefeatDC string //客胜比
		dtcell = dtrow.AddCell()
		dtcell.Value = gamenow.Gamestatistics.GDefeatDC
	}

	//填充内容

	err = file.Save("file.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

//StatisticsGame
func StatisticsGame(game Model.GameNow) {

}
