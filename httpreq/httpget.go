package httpreq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
