package WebCralwer

import (
	"log"
	"myCPforGo/Model"
	"strings"

	"github.com/goinggo/mapstructure"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	tsgutils "github.com/typa01/go-utils"
)

const (
	str_urlTeam = "http://saishi.zgzcw.com/soccer/league/"
)

/*
FindAllLeagueAndCrawlerTeam 找到所有的league并且爬虫保存team
日常task 通过遍历联赛找到所有球队名称存到数据库中
没有做判重，所以用之前需要全部删除DicTeam的数据。
截至2020-03-20 一共有1320条数据
func main(){
	FindAllLeagueAndCrawlerTeam()
}
*/
func FindAllLeagueAndCrawlerTeam() {
	for _, v := range FindAllLeague() {
		CrawlerTeam(v.CodeID) //每一个联赛都去存储数据
	}
}

//SearchForLeague 查找所有的league
//返回 []Model.league 所有的league列表
func FindAllLeague() []Model.League {
	str_sql := "select * from DicLeague"
	var leagues []Model.League
	for _, v := range SearchCom(str_sql) {
		var league Model.League
		if err := mapstructure.Decode(v, &league); err != nil {
			log.Println("转换失败，请查询数据结构是否匹配")
		} else {
			//fmt.Println(league)
			leagues = append(leagues, league)
		}

	}

	return leagues
}

//CrawlerTeam 爬取队伍DIC信息
//code league's code
func CrawlerTeam(code string) {
	str_href := str_urlTeam + code
	//fmt.Println(str_href)
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find("div.tongji_list").Each(func(i int, s *goquery.Selection) {
			s.Find("div").Each(func(i int, ss *goquery.Selection) {
				if ss.Text() == "球队列表" {
					ss.Next().Find("a").Each(func(i int, sss *goquery.Selection) {

						/*存储*/
						code_id, _ := sss.Attr("href")
						//截取字符串
						code_ids := strings.FieldsFunc(code_id, func(c rune) bool {
							if c == '/' {
								return true
							}
							return false
						})
						code_id = code_ids[len(code_ids)-1]
						str_teamname, _ := sss.Attr("title")
						var team Model.Team
						team.TeamName = str_teamname
						team.CodeID = code_id
						team.UUID = tsgutils.UUID()
						SaveComm(team, "DicTeam")

					})
				}
			})
		})
	})
	c.Visit(str_href)
}
