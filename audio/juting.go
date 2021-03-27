package audio

import (
	"CrawlerAudioFiction/httpreq"
	"CrawlerAudioFiction/stx"
	"fmt"
	"log"
	"time"
)

// https://www.yousxs.com/
//https://www.yousxs.com/classify.html?clsid=&pageNum=15&novelName=
// 根据规则解析网页
func GetJuTingPageInfo(url string, stc *stx.ServiceContext) {
	alltitlelist, urllist, err := httpreq.ParseJuTingPageInfo(url)
	if err != nil {
		log.Fatal("解析失败", err)
		return
	}
	for i := 0; i < len(alltitlelist); i ++ {
		alltitle := alltitlelist[i]
		url := "https://www.yousxs.com/" + urllist[i]
		jt, err := httpreq.ParseJuTingOneInfo(url, alltitle)
		if err != nil {
			log.Println("解析报错", url, err)
		}
		time.Sleep(time.Second / 5)
		//insertCode := `INSERT INTO bilibili (id, alltitle, title, url, author, announcer, intro, Intro, videolists) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
		//stc.RW.Lock()
		//result, err := stc.DB.Exec(insertCode, jt.ID, jt.Alltitle, jt.Title, jt.Url, jt.Author, jt.Announcer, jt.Intro, jt.Videolists)
		//if err != nil {
		//	log.Println("插入数据失败：", err)
		//}
		//log.Println("插入成功：", result)
		//stc.RW.Unlock()
	}

}

func CrawlJuTing(stc *stx.ServiceContext) {
	for page := 1; page < 837; page++ {
		url := fmt.Sprintf("https://www.yousxs.com/classify.html?clsid=&pageNum=%d&novelName=", page)
		GetJuTingPageInfo(url, stc)
	}
}