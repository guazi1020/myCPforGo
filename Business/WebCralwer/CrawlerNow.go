package WebCralwer

import (
	"encoding/json"
	"fmt"
	"myCPforGo/Com"
	"myCPforGo/Model"
	"reflect"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//GetEByDate 根据时间来测算每场的E值
func GetEByDate() {
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

			tag := false
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

			if gameNow.GameE != 0 && tag == true && (gameNow.GameE > 1 || gameNow.GameE < 0.923) {

				//计算最近的进球率
				games = append(games, gameNow)
			}

		})
		data, _ := json.Marshal(games)
		fmt.Println(string(data))
	})
	c.Visit("http://cp.zgzcw.com/lottery/jchtplayvsForJsp.action?lotteryId=47&type=jcmini&issue=2020-06-24")
}

//GetEByDate2 get
func GetEByDate2() {
	// c := colly.NewCollector()
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	//e.Request.Visit(e.Attr("href"))
	// 	fmt.Println(e.Attr("href"))
	// })
	// c.Visit("http://live.zgzcw.com/ls/AllData.action?code=201&date=2020-06-16&ajax=true")
	strURL := "http://live.zgzcw.com/ls/AllData.action?code=all&date=2020-01-16&ajax=true"
	params := make(map[string]string)
	params["date"] = "2020-01-16"
	for _, game := range GetWeb(strURL, params) {
		value := reflect.ValueOf(game)
		for i := 0; i < value.NumField(); i++ {
			fmt.Printf("Field %d: %v\n", i, value.Field(i))
		}

	}
}
