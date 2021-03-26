package audio

import (
	"CrawlerAudioFiction/httpreq"
	"CrawlerAudioFiction/model"
	"CrawlerAudioFiction/stx"
	"fmt"
	"log"
)

// 爬取B站的关键词信息去抓取有声小说
// https://api.bilibili.com/x/web-interface/search/all/v2?context=&page=1&order=&keyword=%E5%A4%A9%E5%AE%98%E8%B5%90%E7%A6%8F%E6%9C%89%E5%A3%B0%E5%B0%8F%E8%AF%B4&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&highlight=1&single_column=0
// https://api.bilibili.com/x/web-interface/search/type?context=&page=8&order=&keyword=%E5%A4%A7%E5%A5%89%E6%89%93%E6%9B%B4%E4%BA%BA%20%E6%9C%89%E5%A3%B0%E5%B0%8F%E8%AF%B4&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&search_type=video&highlight=1&single_column=0
// b站抓取规则 首先是抓第一页获取所有页数以及总条目 然后根据合并查找获取所有信息

// 获取第一页拿到最开始的全部page和count
func GetFirstPageInfo(url string) (numPages, numResults, pageSize int) {
	var biliBiliSearchPageAll model.BiliBiliSearchPageAll
	httpreq.ExecGet(&biliBiliSearchPageAll, url, httpreq.HeaderGethttp)
	return biliBiliSearchPageAll.Data.NumPages, biliBiliSearchPageAll.Data.NumResults, biliBiliSearchPageAll.Data.PageSize
}

// AID int `json:"aid"`
// Arcrank string `json:"arcrank"`
// Arcurl string `json:"arcurl"`
// Author string `json:"author"`
// Bvid string `json:"bvid"`
// Description string `json:"description"`
// Duration string `json:"duration"`
// VID int `json:"vid"`
// IsUnionVideo int `json:"is_union_video"`
// MID int `json:"mid"`
// Pic string `json:"pic"`
// Pubdate int `json:"pubdate"`
// Tag string `json:"tag"`
// Title string `json:"title"`
// Type string `json:"type"`
// Typeid string `json:"typeid"`
// Typename string `json:"typename"``

func GetPageDetialInfo(url string, stc *stx.ServiceContext) {
	var biliBiliPageInfo model.BiliBilPageInfo
	httpreq.ExecGet(&biliBiliPageInfo, url, httpreq.HeaderGethttp)

	for _, video := range biliBiliPageInfo.Data.Result {
		video.SetID()
		videoUrl := "https://www.bilibili.com/video/" + video.Bvid
		log.Println(video, videoUrl)
		// insertCode := `INSERT INTO jingting (id, aid, arcrank, arcurl, author, bvid, description, duration, vid, is_union_video, mid, pic, pubdate, tag, title, type, typeid, typename, videourl) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		// stc.RW.Lock()
		// result, err := stc.DB.Exec(insertCode, video.ID, video.Arcrank, video.Arcurl, video.Author, video.Bvid, video.Description, video.Duration, video.VID, video.IsUnionVideo, video.MID, video.Pic, video.Pubdate, video.Tag, video.Title,  video.Type,  video.Typeid, video.Typename, videoUrl)
		// if err != nil {
		// 	log.Println("插入数据失败：", err)
		// }
		// log.Println("插入成功：", result)
		// stc.RW.Unlock()
	}
}

func CrawlBiliBili(keyword string, stc *stx.ServiceContext) {
	// keyword := "%E5%A4%A9%E5%AE%98%E8%B5%90%E7%A6%8F%E6%9C%89%E5%A3%B0%E5%B0%8F%E8%AF%B4"
	word := keyword + "有声小说"
	urlAll := "https://api.bilibili.com/x/web-interface/search/all/v2?context=&page=1&order=&keyword=" + word + "&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&highlight=1&single_column=0"
	numPages, numResults, pageSize := GetFirstPageInfo(urlAll)
	log.Println(keyword, "该书在bilibili的有声小说：", numPages, numResults, pageSize)
	// 根据 numPages, nuResults, PageSize 获取到每一页的信息。在每一页提取到的信息进行保存到数据库
	for i := 1; i*pageSize < numResults; i++ {
		urlPage := fmt.Sprintf("https://api.bilibili.com/x/web-interface/search/type?context=&page=%d&order=&keyword=%s&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&search_type=video&highlight=1&single_column=0", i, word)
		GetPageDetialInfo(urlPage, stc)
	}
}
