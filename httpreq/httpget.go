package httpreq

import (
	"CrawlerAudioFiction/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// 解析网页

// parse http://www.audio699.com/ data
func ParseJingTingPage(url string) (*model.JingTingData, error) {
	var jt = &model.JingTingData{
		Plist: make([]string, 0),
	}
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	// reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)
	// resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("goquery failed parse bofy")
	}
	// 解析出基本信息
	jt.Title = doc.Find(".binfo h1").Text()

	infoPList := doc.Find(".binfo p").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	jt.Class = infoPList[0][9:]
	jt.Author = infoPList[1][9:]
	jt.Announcer = infoPList[2][9:]
	jt.Hits = infoPList[3][9:]
	jt.State = infoPList[4][9:]
	jt.UpdateTime = infoPList[5][9:]
	// 解析出简介
	jt.Intro = doc.Find(".intro p").Text()
	jt.Plist = doc.Find(".playlist ul li a").Map(func(i int, s *goquery.Selection) string {
		url, _ = s.Attr("href")
		return url
	})
	jt.SetID()
	return jt, nil
}

func Gethttp(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return msg, err
}

func ExecGet(res interface{}, host string, get func(url string) ([]byte, error)) error {
	msgByte, err := get(host)
	if err != nil {
		return err
	}
	err = json.Unmarshal(msgByte, res)
	if err != nil {
		return err
	}
	return nil
}

func HeaderGethttp(url string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	// reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return msg, err
}

func ParseXiMaLaYaPage(url string) []string {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	// reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)
	// resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	categoryListName := doc.Find(".category-filter-value-list a").Map(func(i int, s *goquery.Selection) string {
		name, _ := s.Attr("data-code")
		return name
	})
	return categoryListName
}

func ParseJuTingPageInfo(url string)([]string, []string, error) {
	// 定义JuTing网的获取信息结构体
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	// reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)
	// resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	titleList := doc.Find(".panel-body ul li").Map(func(i int, selection *goquery.Selection) string {
		return selection.Find("a").Text()
	})
	urlList := doc.Find(".panel-body ul li").Map(func(i int, selection *goquery.Selection) string {
		url, _ :=  selection.Find("a").Attr("href")
		return url
	})
	return titleList, urlList, nil
}

func ParseJuTingOneInfo(url string, alltitle string)(model.JuTingData, error)  {
	var jt model.JuTingData
	jt.Url = url
	jt.Alltitle = alltitle
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	// reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(reqest)
	// resp, err := http.Get(url)
	if err != nil {
		return jt, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	jt.Title = doc.Find(".col-md-7.col-xs-12.col-sm-7 h3").Text()
	authoranninfo := doc.Find(".col-md-7.col-xs-12.col-sm-7 p").Map(func(i int, selection *goquery.Selection) string {
		return selection.Text()
	})
	jt.Author = authoranninfo[0]
	jt.Announcer = authoranninfo[1]
	jt.Intro = authoranninfo[2]
	jt.VideoList = doc.Find(".panel.panel-success .panel-body div").Map(func(i int, selection *goquery.Selection) string {
		uil, _ := selection.Find("div a").Attr("href")
		return "https://www.yousxs.com/" + uil
	})
	val, err := json.Marshal(jt.VideoList)
	if err != nil {
		log.Println("编码报错")
		return jt, err
	}
	jt.Videolists = string(val)
	jt.SetID()
	return jt, nil
}