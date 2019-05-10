package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func main(){
	lesson74()
}

// json.Unmarshal
func lesson74(){
	// NWから来たjsonをstructに格納
	b := []byte(`{"name":"mike", "age":"20", "nicknames":["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil{
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
	
	//structをjsonに変換
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}

// json.Marshal をオーバーライドする
func (p Person) MarshalJSON() ([]byte, error){
	//a := &struct{Name string}{Name: "test"}
	v, err := json.Marshal(&struct{
			Name string
		}{
			Name: "Mr." + p.Name,
		})
	return v, err
}

type Person struct{
	// ``で囲むとjson変換時のキーを指定できる
	Name       string  `json:"name"`  //`json:name,omitempty` で空要素のときに要素を非表示にできる
	Age        int     `json:"age,string"`
	Nicknames []string `json:"nicknames"`
	T          T       `json:T,omitempty`
}

type T struct{}

// http
func lesson73(){
	/* シンプルなGETメソッド
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	*/
	
	//URLの形式チェック
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	
	//リクエストを投げる
	req, _ := http.NewRequest("GET", endpoint, nil)
	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("IF-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()
	
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}





