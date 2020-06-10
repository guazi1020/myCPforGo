package WebCralwer

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//GetEByDate 根据时间来测算每场的E值
func GetEByDate() {
	//1.根据时间爬取当前页面
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		dom.Find("table.mb").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				//var teamname []string
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					fmt.Println(s.Find("a").Text())
				})
				//fmt.Println(s.Attr("m"))
			})
		})
	})
	c.Visit("http://cp.zgzcw.com/lottery/jchtplayvsForJsp.action?lotteryId=47&type=jcmini&issue=2020-06-10")
}
