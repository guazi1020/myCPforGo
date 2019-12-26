package WebCralwer

import (
	"context"
	"log"
	"myCPforGo/Com"
	"myCPforGo/Model"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func SaveWeb() {
	d := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer func() { cancel() }()
	ch := make(chan int)
	games := GetWeb()
	log.Println("开始工作")
	for _, game := range games {
		go SaveOneGameInfo(game, ctx, ch)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println(<-ch)
			log.Println("工作完了")
			return
		case <-ch:
			log.Println("通道改变了")
		}
	}
}

func GetWeb() []Model.Game {
	games := []Model.Game{}
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find("em.input-xh").Each(func(i int, s *goquery.Selection) {
			var game Model.Game
			game.Gnumber = Com.RemoveBlank(s.Text())
			games = append(games, game)

		})
		dom.Find("span.sptr").Children().Each(func(i int, s *goquery.Selection) {
			_, bl := s.Attr("href")
			if bl {
				//fmt.Println("编号", i)
				games[(i-3)/5].GhomeName = s.Text()
			}
			strClass, result := s.Attr("class")
			if result == true {
				switch strClass {
				case "hongpai": //比赛红牌
					games[i/5].GredQuantities = Com.RemoveBlank(s.Text())
					break
				case "paim": //主队排名
					games[(i-2)/5].GhomeRank = Com.RemoveCharacter(Com.RemoveCharacter(s.Text(), "", "]"), "", "[")
				case "rq": //主队让球
					games[(i-4)/5].GletCount = Com.RemoveCharacter(Com.RemoveCharacter(s.Text(), "", ")"), "", "(")
				default:
					break
				}
			}
		})
		//比赛结果
		dom.Find("span.boldbf").Each(func(i int, s *goquery.Selection) {
			//fmt.Println(i,s.Text())
			games[i].GresultScore = Com.RemoveBlank(s.Text())
		})
		//赛制
		dom.Find("body>span[style]").Each(func(i int, s *goquery.Selection) {
			//fmt.Println(i,s.Text())
			games[i].Gleague = Com.RemoveBlank(s.Text())
			//fmt.Println(s.Html())
			//s.Find
		})

		dom.Find("span.sptl").Children().Each(func(i int, s *goquery.Selection) {
			_, bl := s.Attr("href")
			if bl {
				//fmt.Println("编号", i)
				games[(i)/4].GguestName = s.Text()
			}
			strClass, result := s.Attr("class")
			if result == true {
				switch strClass {
				case "hongpai": //客队比赛红牌
					// fmt.Println("编号", i)
					games[(i-3)/4].GredQuantitlesGuest = Com.RemoveBlank(s.Text())
					break
				case "paim": //客队排名
					games[(i-1)/4].GguestRank = Com.RemoveCharacter(Com.RemoveCharacter(s.Text(), "", "]"), "", "[")
				default:
					break
				}
			}
		})
		//半场比分
		dom.Find("span.bcbf").Each(func(i int, s *goquery.Selection) {
			games[i].GresultHalfScore = Com.RemoveBlank(s.Text())
		})
		//最终结果
		dom.Find("strong.f_sf").Each(func(i int, s *goquery.Selection) {
			games[i].Gresult = Com.RemoveBlank(s.Text())
		})
		//是否完成比赛
		dom.Find("strong").Each(func(i int, s *goquery.Selection) {
			strclass, _ := s.Attr("class")
			if strclass == "red" {
				games[i/8].GIsfinish = Com.RemoveBlank(s.Text())
			}
		})
		//比赛欧赔
		dom.Find("div.oupei").Each(func(i int, s *goquery.Selection) {
			s.Find("span").Each(func(j int, t *goquery.Selection) {
				switch j {
				case 0:
					games[i].GspWin = Com.RemoveBlank(t.Text())
					break
				case 1:
					//fmt.Println(i,t.Text())
					games[i].GspTie = Com.RemoveBlank(t.Text())
					break
				case 2:
					games[i].GspDefeat = Com.RemoveBlank(t.Text())
					break
				default:
					break
				}
			})
		})
		// for index, value := range games {
		// 	fmt.Println(index,value)
		// }

	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		//log.Println(e.ChildAttr(".matchDate", "date"))
		//	log.Println(e)
	})
	c.Visit("http://live.zgzcw.com/ls/AllData.action?code=201&date=2019-11-23&ajax=true")
	return games
}
