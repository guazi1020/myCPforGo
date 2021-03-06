/**
**	爬数据
**	http://live.zgzcw.com/ls/AllData.action
**/

package WebCralwer

import (
	"context"
	"fmt"
	"log"
	"myCPforGo/Com"
	"myCPforGo/Com/baseMethod"
	"myCPforGo/Model"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const (
	timeLayout = "2006-01-02 15:04:05"
	dateLayout = "2006-01-02"
)

//SaveWebByDate 根据开始时间和结束时间爬站点的数据，存入数据库中
//beginDate 开始时间,为空默认为只爬endDate一天
//endDate 结束时间，为空默认为当前一天
func SaveWebByDate(beginDate string, endDate string, params map[string]string) bool {

	//条件判断，判断当前日期的问题
	if endDate == "" {
		endDate = time.Now().Format(dateLayout)
	}
	if beginDate == "" {
		beginDate = time.Now().Format(dateLayout)
	}

	loc, _ := time.LoadLocation("Local")
	thebeginDate, _ := time.ParseInLocation(dateLayout, beginDate, loc)
	theendDate, _ := time.ParseInLocation(dateLayout, endDate, loc)

	beginSr := thebeginDate.Unix()
	//end_sr := theendDate.AddDate(0, 0, -1).Unix()
	endSr := theendDate.AddDate(0, 0, -1).Unix()

	if endSr-beginSr < 0 {
		return false //如果最晚时间早于开始时间，结束
	}
	fmt.Println((endSr - beginSr) / 86400)
	for i := 0; i <= int((endSr-beginSr)/86400); i++ {

		//i-1 yesteday,no today
		params["date"] = thebeginDate.AddDate(0, 0, i).Format(dateLayout)
		tableName := "game"
		primaryField := "Gyear"
		switch params["code"] {
		case "all":
			{
				tableName = "GameAllBasic"
				primaryField = "GAdate"
			}
		}
		if IsOnly(params["date"], tableName, primaryField) {
			fmt.Println("已经更新过了")
			continue
		}
		SaveWeb(params)
	}

	//clearRepeatInfo
	ClearRepeatInfo()
	return true
}

//SaveWeb 根据参数开始工作
// params := make(map[string]string)
// params["code"] = "201"
// params["ajax"] = "true"
// params["date"]="2020-02-01"
func SaveWeb(params map[string]string) {
	d := time.Now().Add(5000 * time.Millisecond) //5秒最大限度
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer func() { cancel() }()
	strHref := CompositionURL("http://live.zgzcw.com/ls/AllData.action", params) //组装href

	switch params["code"] {
	case "201":
		games := GetWeb(strHref, params)
		log.Println("开始工作", strHref)
		//fmt.Println(games)
		for key, game := range games {
			// fmt.Println(game)
			// fmt.Println(key)
			go SaveOneGameInfo(game, ctx, key) //开协程

		}
	case "all":
		gameAlls := GetWebToGames(strHref, params)

		fmt.Println("时间：", params["date"])
		//fmt.Println(gameAlls)
		for key, gameAll := range gameAlls {
			go SaveOneGameAllBaiscInfo(gameAll, ctx, key) //开协程
		}
		///fmt.Println(gameAlls)
	default:
		break
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("工作完了")
			return
		}
	}
}

//GetWeb 获取页面数量并对象化
//面向参数为201
//http://live.zgzcw.com/ls/AllData.action?code=all&date=2020-01-16&ajax=true
func GetWeb(str_href string, paras map[string]string) []Model.Game {

	games := []Model.Game{}
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))

		dom.Find("em.input-xh").Each(func(i int, s *goquery.Selection) {

			var game Model.Game
			game.Gnumber = Com.RemoveBlank(s.Text())
			games = append(games, game)

		})

		if len(games) == 0 {
			fmt.Println("切片为零")
			dom.Find("input[name='order']").Each(func(i int, s *goquery.Selection) {
				//fmt.Println(i)
				var game Model.Game
				game.Gnumber = Com.RemoveBlank(s.Text())
				games = append(games, game)

			})
		}
		//fmt.Print("bbegin")
		dom.Find("span.sptr").Children().Each(func(i int, s *goquery.Selection) {

			_, bl := s.Attr("href")
			if bl {

				games[(i-3)/5].GhomeName = s.Text()
			}
			strClass, result := s.Attr("class")
			if result == true {
				switch strClass {
				case "hongpai": //比赛红牌

					games[i/5].GredQuantities = Com.RemoveBlank(s.Text())
					break
				case "paim": //主队排名
					//	fmt.Println(i)
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
		dom.Find("body").Each(func(i int, s *goquery.Selection) {
			//remove space and make []string
			strs_ := strings.FieldsFunc(s.Text(), unicode.IsSpace)
			y := 0
			x := 0

			// gst := make(map[int]map[int]string)
			gst_item := make(map[int]string)
			for _, str := range strs_ {
				//fmt.Println(str)
				gst_item[x] = str
				x++
				if str == "析" {
					x = 0
					games[y].Gdata = gst_item[3]
					games[y].Gtime = gst_item[4]
					games[y].Gleaguenumber = gst_item[2]
					games[y].Gyear = paras["date"]
					y++
				}
				//fmt.Println(gst)
			}
			//fmt.Println(gst)
			// for key, _ := range gst {
			// 	//games[key].Gleaguenumber = value[3]
			// 	fmt.Println(key)
			// }

		})
		// for index, value := range games {
		// 	fmt.Println(index,value)
		// }

	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		//log.Println(e.ChildAttr(".matchDate", "date"))
		//	log.Println(e)
	})
	//c.Visit("http://live.zgzcw.com/ls/AllData.action?code=201&date=2019-11-23&ajax=true")
	c.Visit(str_href)
	return games
}

//CompositionURL 拆解组合URL
func CompositionURL(head string, params map[string]string) string {
	//1.拆params
	var str_params string
	for key, value := range params {
		str_params += key + "=" + value + "&"
	}
	if str_params != "" {
		str_params = strings.TrimRight(str_params, "&")
	}
	//2.组装url
	return head + "?" + str_params
}

//GetWebToGames 组装全量数据对象
func GetWebToGames(strHref string, params map[string]string) []Model.GameAllBasic {
	games := []Model.GameAllBasic{}
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		//fmt.Println(dom.Html())
		dom.Find("body").Each(func(i int, hs *goquery.Selection) {
			hs.Find("input[name='order']").Each(func(i int, s *goquery.Selection) {
				//fmt.Println(i)
				var game Model.GameAllBasic
				//fmt.Println(s.Next().Next().Text())
				//game.GAnumber = Com.RemoveBlank(s.Text())

				//Gleague
				game.GAleague = s.Next().Text()
				s.Next().Remove()
				//GIsfinish
				game.GAIsFinished = s.Next().Text()
				s.Next().Remove()
				//HomeRank
				game.GAHomeRank = s.Next().Find("em.paim").Text()
				if len(game.GAHomeRank) > 1 {
					game.GAHomeRank = game.GAHomeRank[1 : len(game.GAHomeRank)-1]
				}
				//s.Next().Remove()
				//HomeName
				game.GAHomeName = s.Next().Find("a").Text()
				s.Next().Remove()
				//resultsource
				game.GAresultScore = s.Next().Text()
				s.Next().Remove()
				//GuestName
				game.GAGuestName = s.Next().Find("a").Text()
				//GuestRank
				game.GAGuestRank = s.Next().Find("em").Text()
				if len(game.GAGuestRank) > 1 {
					game.GAGuestRank = game.GAGuestRank[1 : len(game.GAGuestRank)-1]
				}
				s.Next().Remove()

				game.GAresultHalf = s.Next().Text()
				s.Next().Remove()

				game.GAspWin = s.Next().Find("span").Eq(0).Text()
				game.GAspTie = s.Next().Find("span").Eq(1).Text()
				game.GAspDefeat = s.Next().Find("span").Eq(2).Text()

				gguestRank, _ := strconv.ParseFloat(game.GAGuestRank, 64)
				ghomeRank, _ := strconv.ParseFloat(game.GAHomeRank, 64)
				gspwin, _ := strconv.ParseFloat(game.GAspWin, 64)
				//计算E值
				game.GAE = baseMethod.ChangeNumber(Calculate_E(gguestRank-ghomeRank, gspwin), 3)

				game.GAresult = baseMethod.CalculateGameResult(game.GAresultScore)

				game.GADate = params["date"]
				//判定是否符合要求
				games = append(games, game)
			})

			/*单独填充轮数(round)*/
			strsource := Com.RemoveBlank(hs.Text())
			//fmt.Println(strsource)
			num := 0
			index := 0
			tempStrings := []string{}
			for k, item := range strings.Fields(strsource) {
				//fmt.Println(string(item))
				if k != 0 {
					tempStrings = append(tempStrings, item)
					if item == "欧亚析" {
						tag := 0
						rs := []rune(tempStrings[0])
						if string(rs[0:1]) == "周" {
							tag = 1
						}

						games[index].GARound = tempStrings[tag]
						num = 0
						index++
						tempStrings = []string{}
					} else {
						num++
					}
				}
			}

		})

	})
	c.Visit(strHref)
	return games
}

//GetInfoByE 根据E来放回
func GetInfoByE(strE string) {

}
