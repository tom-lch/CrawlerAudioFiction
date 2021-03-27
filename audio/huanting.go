package audio

import (
	"CrawlerAudioFiction/stx"

	"fmt"
)

func GetHuanTingPage(url string, stc *stx.ServiceContext) {

}

// http://www.ting89.com/topiclist/quanben.html
// http://www.ting89.com/
func CrawlHuanTing(stc *stx.ServiceContext) {
	for page := 1; page < 508; page++ {

		url := fmt.Sprintf("http://www.ting89.com/topiclist/quanben-%d.html", page)
		if page == 1 {
			url = "http://www.ting89.com/topiclist/quanben.html"
		}
		GetHuanTingPage(url, stc)
	}
}
