package audio

import (
	"CrawlerAudioFiction/httpreq"
	"CrawlerAudioFiction/model"
	"CrawlerAudioFiction/stx"

	"fmt"
	"log"
)

// https://www.ximalaya.com/revision/category/queryCategoryPageAlbums?category=youshengshu&subcategory=wenxue&meta=&sort=0&page=1&perPage=30

// 喜马拉雅规则比较复杂，只能爬取部分的数据，并不能获得全部的有效数据
// 1 https://www.ximalaya.com/youshengshu/tongshu/ 抓取该页面小的有声书的所有分类

func GetAllCategory(url string, stc *stx.ServiceContext) {
	CategoryNames := httpreq.ParseXiMaLaYaPage(url)
	for i := 0; i < len(CategoryNames); i++ {
		name := CategoryNames[i]
		for page := 1; page < 11; page++ {
			url := fmt.Sprintf("https://www.ximalaya.com/revision/category/queryCategoryPageAlbums?category=youshengshu&subcategory=%s&meta=&sort=0&page=%d&perPage=30", name, page)
			GetDetailPageInfo(url, stc)
			// time.Sleep(time.Second * 5)
		}
	}
}

func GetDetailPageInfo(url string, stc *stx.ServiceContext) {
	var xm model.XimalayaPageInfo
	err := httpreq.ExecGet(&xm, url, httpreq.HeaderGethttp)
	if err != nil {
		log.Println(err)
	}
	log.Println(xm)
	for _, album := range xm.Data.Albums {
		album.SetID()
		albumUrl := "https://www.ximalaya.com/youshengshu/" + album.Link
		log.Println(album, albumUrl)
		// insertCode := `INSERT INTO ximalaya (id, albumid, albumsubscript, anchorname, coverpath, isfinished, ispaid, link, title, uid) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		// stx.RW.Lock()
		// result, err := stx.DB.Exec(insertCode, album.ID, album.AlbumId, album.AlbumSubscript, album.AnchorName, album.CoverPath, album.IsFinished, album.IsPaid, albumUrl, album.Title, album.Uid)
		// if err != nil {
		// 	log.Println("插入数据失败：", err)
		// }
		// log.Println("插入成功：", result)
		// stx.RW.Unlock()
	}
}

// 抓取web浏览器上的分类数据
func CrawlXimalaya(stc *stx.ServiceContext) {
	url := "https://www.ximalaya.com/youshengshu/tongshu/"
	GetAllCategory(url, stc)
}
