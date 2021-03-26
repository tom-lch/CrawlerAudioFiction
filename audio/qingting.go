package audio

import (
	"CrawlerAudioFiction/httpreq"
	"CrawlerAudioFiction/model"
	"CrawlerAudioFiction/stx"
	"fmt"
	"log"
	"strconv"
)

// https://i.qingting.fm/capi/neo-channel-filter?category=521&attrs=3290&curpage=1 女生
// https://i.qingting.fm/capi/neo-channel-filter?category=521&attrs=3289&curpage=1 男生

func GetQingTingPageInfo(url string, stc *stx.ServiceContext) {
	var qt model.QingTingPageInfo
	err := httpreq.ExecGet(&qt, url, httpreq.HeaderGethttp)
	if err != nil {
		log.Println(err)
	}
	log.Println(qt)
	for _, Channel := range qt.Data.Channels {
		Channel.SetID()
		Channelurl := fmt.Sprintf("https://www.qingting.fm/channels/%d", Channel.QtID)
		log.Println(Channel, Channelurl)
		// insertCode := `INSERT INTO qingitng (id, cover, description, qtid, playcount, title, type, update_time, channelurl) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
		// stx.RW.Lock()
		// result, err := stx.DB.Exec(insertCode, Channel.ID, Channel.Cover, Channel.Description, Channel.QtID, Channel.Playcount, Channel.Title, Channel.Type, Channel.UpdateTime, Channelurl)
		// if err != nil {
		// 	log.Println("插入数据失败：", err)
		// }
		// log.Println("插入成功：", result)
		// stx.RW.Unlock()
	}
}

func CrawlQingTing(stc *stx.ServiceContext) {
	//根据男生 女生来搜索蜻蜓
	for i := 3289; i <= 3290; i++ {
		url := fmt.Sprintf("https://i.qingting.fm/capi/neo-channel-filter?category=521&attrs=%d&curpage=", i)
		for page := 1; page < 76; page++ {
			url = url + strconv.Itoa(page)
			GetQingTingPageInfo(url, stc)
		}
	}
}
