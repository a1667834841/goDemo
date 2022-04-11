package smzdm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type result struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Data      data   `json:"data"`
}

type data struct {
	Rows  []product `json:"rows"`
	Total int       `json:"total"`
}

type product struct {
	ArticleTitle  string `json:"article_title"`
	ArticlePrice  int    `json:"article_price"`
	ArticleWorthy int    `json:"article_worthy"`
}

func GetGoods() {
	params := url.Values{}
	Url, err := url.Parse("https://api.smzdm.com/v1/home/articles_new")
	if err != nil {
		return
	}
	params.Set("f", "wxapp")
	params.Set("wxapp", "wxapp")
	params.Set("offset", "0")
	params.Set("limit", "20")

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath) //
	resp, err := http.Get(urlPath)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)

}
