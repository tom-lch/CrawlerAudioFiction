package httpreq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func ExecPost(rep, res interface{}, host string) error {
	jsonStr, err := json.Marshal(rep)
	if err != nil {
		return err
	}
	msgByte, err := httppost(host, jsonStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal(msgByte, res)
	if err != nil {
		return err
	}
	return nil
}

func httppost(host string, jsonStr []byte) ([]byte, error) {
	// hc := http.Client{}
	// req, err := http.NewRequest("POST", host, bytes.NewBuffer(jsonStr))

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// resp, err := hc.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	resp, err := http.Post(host,
		"application/x-www-form-urlencoded",
		strings.NewReader(string(jsonStr)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	msgByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return msgByte, nil
}
