package main

import (
	"CrawlerAudioFiction/httpreq"
	// "Libs/Rsystem"
	"fmt"
	"log"
	"strconv"
)

//爬取静听网上的有声读物信息  http://www.audio699.com/
// 爬取规则很简单，根据url规则进行爬取 http://www.audio699.com/book/5410.html    1 - 5410

func GetEveryPage(start, end int, url string) {
	for i := start; i < end; i++ {
		index := strconv.Itoa(i)
		host := url + index + ".html"
		fmt.Println(host)
		msg, err := httpreq.Gethttp(host)
		if err != nil {
			log.Println(host, err)
		}
		ParsePage(msg)
	}
}

// 解析网页
func ParsePage(info []byte) {
	fmt.Println(string(info))
	re := 
}

func main() {
	fmt.Println("start")
	GetEveryPage(2345, 2346, "http://www.audio699.com/book/")
}