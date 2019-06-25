package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
	"log"
	"regexp"
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
	fmt.Printf("## " + filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Printf("No file.... " + filename)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// 毎回rendertemplateで描画されているものをキャッシングする
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
//	t, _ := template.ParseFiles(tmpl + ".html")
//	t.Execute(w, p)
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	// rにはアクセスした情報（httpリクエストの内容）が含まれる
	// wにはレスポンスに加える内容を付加する
	p, err := loadPage(title)
		if err != nil{
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string){
	fmt.Println(title)
	p, err := loadPage(title)
	
	if err != nil{
		http.Redirect(w, r, "/view/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string){
	body := r.FormValue("body")
	p := &Page{Title:title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main(){
	p1 := &Page{Title: "view", Body: []byte("This is a sample Page.")}
	p1.save()
//	
//	p2, _ := loadPage(p1.Title)
//	fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8090", nil))
}