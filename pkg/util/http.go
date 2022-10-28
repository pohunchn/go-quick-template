package util

import (
	"go-quick-template/pkg/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ReqParams struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Post(url string, params url.Values) *ReqParams {
	resp, err := http.PostForm(url, params)
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			resp.Body.Close()
			retData := &ReqParams{}
			if err := json.Unmarshal(body, retData); err != nil {
				return nil
			}
			return retData
		}
		return nil
	}
	return nil
}
