package audio

import (
	"CrawlerAudioFiction/stx"

	"fmt"
)

func GetHuanTingPage(url string, stc *stx.ServiceContext) {
	
}

// http://www.ting89.com/
func CrawlHuanTing(stc *stx.ServiceContext) {
	for page := 1; page < 419; page++ {
		url := fmt.Sprintf("https://www.yousxs.com/rankinglist.html?type=1&pageNum=%d", page)
		GetHuanTingPage(url, stc)
	}
}
