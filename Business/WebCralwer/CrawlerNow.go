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
					fmt.Println(i)
				})
				//fmt.Println(s.Attr("m"))
			})
		})
	})
	c.Visit("http://cp.zgzcw.com/lottery/jchtplayvsForJsp.action?lotteryId=47&type=jcmini&issue=2020-06-10")
}

//GetEByDate2 get
func GetEByDate2() {
	// c := colly.NewCollector()
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	//e.Request.Visit(e.Attr("href"))
	// 	fmt.Println(e.Attr("href"))
	// })
	// c.Visit("http://live.zgzcw.com/ls/AllData.action?code=201&date=2020-06-16&ajax=true")
	strURL := "http://live.zgzcw.com/ls/AllData.action?code=201&date=2020-01-16&ajax=true"
	params := make(map[string]string)
	params["date"] = "2020-01-16"
	fmt.Println(GetWeb(strURL, params))
}
