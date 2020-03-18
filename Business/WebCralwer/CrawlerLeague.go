package WebCralwer

import (
	"myCPforGo/Com"
	"myCPforGo/Model"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const (
	str_url = "http://saishi.zgzcw.com/soccer/"
)

//CrawlerLeague 爬字典league
func CrawlerLeague() {

	c := colly.NewCollector()
	//爬到之后的操作
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find(".ls").Children().Each(func(i int, s *goquery.Selection) {
			str_c, _ := s.Find("div.lstitle").Html() //国家

			s.Find("div.kuang").Each(func(i int, ss *goquery.Selection) {
				ss.Find("a").Each(func(i int, sss *goquery.Selection) { //查找详细
					code_id, _ := sss.Attr("href") //league 编码
					code_name := sss.Text()        //league 名称
					//	log.Println(ss.Html())
					//截取字符串
					code_ids := strings.FieldsFunc(code_id, func(c rune) bool {
						if c == '/' {
							return true
						}
						return false
					})

					code_id = code_ids[len(code_ids)-1]
					var league Model.League
					league.CountryName = Com.RemoveBlank(str_c)
					league.LeagueName = Com.RemoveBlank(code_name)
					league.CodeID = code_id

					SaveLeague(league) //保存就好了
				})
			})
		})
	})

	c.Visit(str_url)
}
