package WebCralwer

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const (
	str_url = "http://saishi.zgzcw.com/soccer/"
)

func CrawlerLeague() {

	//var league Model.League
	//1.找到目标地址
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find(".ls").Children().Each(func(i int, s *goquery.Selection) {
			str_c, _ := s.Find("div.lstitle").Html() //国家
			log.Println(str_c)

			s.Find("div.kuang").Each(func(i int, ss *goquery.Selection) {
				ss.Find("a").Each(func(i int, sss *goquery.Selection) {
					code_id, _ := sss.Attr("href")  //league 编码
					code_name := ss.Text()          //league 名称
					log.Println(code_id, code_name) //

				})
			})
		})
	})

	c.Visit(str_url)
	//2.分析其内容

	//3.数据库操作

}
