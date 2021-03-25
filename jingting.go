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
		jt, err := httpreq.ParsePage(host)
		if err != nil {
			log.Println(host, err)
		}
		// 存入数据库
		fmt.Println(jt)
	}
}



func main() {
	fmt.Println("start")
	GetEveryPage(1, 5411, "http://www.audio699.com/book/")
}
