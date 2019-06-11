package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"log"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	// rにはアクセスした情報（httpリクエストの内容）が含まれる
	// wにはレスポンスに加える内容を付加する
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main(){
//	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
//	p1.save()
//	
//	p2, _ := loadPage(p1.Title)
//	fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}