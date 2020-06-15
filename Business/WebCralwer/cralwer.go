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
	"myCPforGo/Model"
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
func SaveWebByDate(beginDate string, endDate string, params map[string]string) {

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
		return //如果最晚时间早于开始时间，结束
	}

	for i := 0; i <= int((endSr-beginSr)/86400); i++ {
		//i-1 yesteday,no today
		params["date"] = thebeginDate.AddDate(0, 0, i).Format(dateLayout)
		if IsOnly(params["date"]) {
			fmt.Println("true")
			continue
		}

		SaveWeb(params)
	}
}

//SaveWeb 根据参数开始工作
func SaveWeb(params map[string]string) {
	d := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer func() { cancel() }()

	str_href := CompositionURL("http://live.zgzcw.com/ls/AllData.action", params)
	games := GetWeb(str_href, params)
	log.Println("开始工作", str_href)
	//fmt.Println(games)
	for key, game := range games {
		//fmt.Println(game)
		//fmt.Println(key)
		go SaveOneGameInfo(game, ctx, key) //开协程

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
				fmt.Println(i)
				var game Model.Game
				game.Gnumber = Com.RemoveBlank(s.Text())
				games = append(games, game)

			})
		}
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
					fmt.Println(i)
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
