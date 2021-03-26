package main

import (
	"CrawlerAudioFiction/audio"
	"CrawlerAudioFiction/stx"
)

func main() {
	stc := stx.NewServiceContext("conf.yaml")
	// audio.CrawlJingTing(stc)
	// audio.CrawlBiliBili("大奉打更人", stc)
	// audio.CrawlXimalaya(stc)
	audio.CrawlQingTing(stc)
}
