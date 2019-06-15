package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
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
	fmt.Printf(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Printf("No file.... " + filename)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	// rにはアクセスした情報（httpリクエストの内容）が含まれる
	// wにはレスポンスに加える内容を付加する
//	title := r.URL.Path[len("/view/"):]
	title := "view"
	fmt.Println(title)
	p, _ := loadPage(title)
		if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	fmt.Println(title)
	p, err := loadPage(title)
	
	if err != nil{
		http.Redirect(w, r, "/view/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := $Page{Title:title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main(){
	p1 := &Page{Title: "view", Body: []byte("This is a sample Page.")}
	p1.save()
//	
//	p2, _ := loadPage(p1.Title)
//	fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}