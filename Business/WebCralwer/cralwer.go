package WebCralwer

import (
	"fmt"
	"log"
	"myCP/Com"
	"myCP/Model"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func GetWeb() {
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
				case"rq": //主队让球
					games[(i-4)/5].GletCount= Com.RemoveCharacter(Com.RemoveCharacter(s.Text(), "", ")"), "", "(")
				default:
					break
				}
			}
		})

		// dom.Find("span").Each(func(i int,s *goquery.Selection){
		// 	fmt.Println(s.Text())
		// })
		// dom.Find("span.sptr a").Each(func(i int, s *goquery.Selection) {
		// 	fmt.Println(s.Text())
		// })
		for _, value := range games {
			fmt.Println(value)
		}

		// //fmt.Println(dom.Html())
		// dom.Find("em").Siblings().Each(func(i int, s *goquery.Selection) {
		// 	class, _ := s.Attr("class")
		// 	if class == "input-xh" {
		// 		fmt.Println(class)
		// 	}
		// 	fmt.Println(s.Find("a").Html())
		// 	// fmt.Println(s.Html())
		// 	// var game Model.Game

		// 	// switch class {
		// 	// case "input-xh":
		// 	// 	game.Gnumber = s.Text()
		// 	// 	break
		// 	// case "paim":
		// 	// 	fmt.Println(s.Text())
		// 	// 	game.GhomeRank = Com.RemoveBlank(s.Text())
		// 	// 	break
		// 	// default:
		// 	// 	break
		// 	// }
		// 	//game.SaveGametoDB()
		// })
	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		//log.Println(e.ChildAttr(".matchDate", "date"))
		log.Println(e)
	})
	c.Visit("http://live.zgzcw.com/ls/AllData.action?code=201&date=2019-11-23&ajax=true")
}
