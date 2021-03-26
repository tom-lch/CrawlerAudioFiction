package audio

import (
	"CrawlerAudioFiction/httpreq"
	"CrawlerAudioFiction/stx"

	// "Libs/Rsystem"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

//爬取静听网上的有声读物信息  http://www.audio699.com/
// 爬取规则很简单，根据url规则进行爬取 http://www.audio699.com/book/5410.html    1 - 5410

func GetEveryPage(start, end int, url string, stc *stx.ServiceContext) {
	for i := start; i < end; i++ {
		index := strconv.Itoa(i)
		bookUrl := url + index + ".html"
		fmt.Println(bookUrl)
		jt, err := httpreq.ParseJingTingPage(bookUrl)
		if err != nil {
			log.Println(bookUrl, err)
		}
		// 存入数据库
		data, err := json.Marshal(jt.Plist)
		if err != nil {
			log.Println("序列化切片报错：", err)
		}
		jt.Plists = string(data)
		jt.Host = " http://www.audio699.com/"
		jt.Url = bookUrl
		log.Println(jt)
		// insertCode := `
		// INSERT INTO jingting (id, title, class, author, hists, state, update_time, intro, url, host, plist) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		// `
		// stc.RW.Lock()
		// res, err := stc.DB.Exec(insertCode, jt.ID, jt.Title, jt.Class, jt.Author, jt.Hists, jt.State, jt.UpdateTime, jt.Intro, jt.Url, jt.Host, jt.Plists)
		// if err != nil {
		// 	log.Println("插入数据失败：", err)
		// }
		// log.Println("插入成功：", res)
		// stc.RW.Unlock()
	}
}

func CrawlerJingTing(stc *stx.ServiceContext) {
	fmt.Println("start")
	// GetEveryPage(1, 5411, "http://www.audio699.com/book/", stc)
	GetEveryPage(2586, 2587, "http://www.audio699.com/book/", stc)
}
