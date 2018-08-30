package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://shfwauth.montage-tech.com/web_auth/index.html"

	payload := strings.NewReader("edit_cookies=&admin_id=juxu&admin_pw=o5%26lejrffz")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Origin", "https://shfwauth.montage-tech.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Referer", "https://shfwauth.montage-tech.com/web_auth/index.html")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("Postman-Token", "83ed297f-6fb5-4789-a806-2ba45bbd0ef9")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
