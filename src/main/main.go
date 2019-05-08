package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main(){
	lesson73()
}

// http
func lesson73(){
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	
	//URLの形式チェック
	base, err := url.Parse("http:// example.com")
	fmt.Println(base, err)
}





