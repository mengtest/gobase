package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

// Cc 爬虫测试框架
func Cc() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{"quce.cdn.woquhudong.cn", "m.quce001.com"}
	c.OnRequest(func(r *colly.Request) {
		//log.Println("正在拜访一个链接:", r.AbsoluteURL(r.URL.String()))
	})

	c.OnError(func(r *colly.Response, err error) {
		//log.Println(fmt.Sprintf("访问%s链接出现了一个错误:", r.Request.URL), err)
	})

	c.OnResponse(func(r *colly.Response) {
		//log.Println("完成了一个链接的访问 ===> ", r.Request.URL)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//log.Println("发现了一个新的连接,发起访问:", e.Request.AbsoluteURL(e.Attr("href")))
		e.Request.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnHTML("img", func(e *colly.HTMLElement) {
		srcResp, srcErr := http.Get(e.Request.AbsoluteURL(e.Attr("src")))
		if srcErr == nil {
			data, _ := ioutil.ReadAll(srcResp.Body)
			arr := strings.Split(srcResp.Request.URL.RequestURI(), "?")
			fileName := strings.Replace(arr[0], "/", "_", -1)
			ioutil.WriteFile("./image/"+fileName, data, 0666)
			srcResp.Body.Close()
		} else {
			log.Println(srcErr.Error())
		}
		datasrcResp, datasrcErr := http.Get(e.Request.AbsoluteURL(e.Attr("data-src")))
		if datasrcErr == nil {
			data, _ := ioutil.ReadAll(datasrcResp.Body)
			arr := strings.Split(datasrcResp.Request.URL.RequestURI(), "?")
			fileName := strings.Replace(arr[0], "/", "_", -1)
			//fmt.Println(arr[0])
			err := ioutil.WriteFile("./image/"+fileName, data, 0666)
			if err != nil {
				log.Println(err.Error())
			}
			datasrcResp.Body.Close()
		} else {
			log.Println(datasrcErr.Error())
		}
	})
	c.OnHTML("div[data-href]", func(e *colly.HTMLElement) {
		log.Println(e.Request.AbsoluteURL(e.Attr("data-href")))
		e.Request.Visit(e.Request.AbsoluteURL(e.Attr("data-href")))
	})

	c.OnScraped(func(r *colly.Response) {
		//log.Println("爬取成功了一个链接:", r.Request.URL)
	})
	c.Visit("http://m.quce001.com/index.php/wetest/constell/cardindex")
}
