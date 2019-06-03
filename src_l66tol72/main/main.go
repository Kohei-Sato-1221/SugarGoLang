package main

import (
	"github.com/markcheno/go-quote"
	alias "github.com/markcheno/go-talib"
	"fmt"
	"mylib" // 他ファイルのGOファイルを呼ぶ方法
	"time"
	"regexp"
	"sort"
	"context"
	"io/ioutil"
	"log"
	"bytes"
)

func main(){
	lesson72()
}


//ioutil
func lesson72(){
	content, err := ioutil.ReadFile("/Users/koheisato/eclipse-workspace/SugarGoLang/src/main/main.go")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(content))
	
	/* ファイルの書き込み
	if err := ioutil.WriteFile("/Users/koheisato/eclipse-workspace/SugarGoLang/src/main/testio.go", content, 0666); err != nil{
		log.Fatalln(err)
	}
	*/
	
	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))
		
}

// context
func lesson71(){
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
	go longProcess(ctx, ch)
	
	CTXLOOP:
		for {
			select {
				case <- ctx.Done():
					fmt.Println(ctx.Err())
					break CTXLOOP
				case <- ch:
					fmt.Println("success")
					break CTXLOOP
			}
		}
	fmt.Println("################")	
}

func longProcess(ctx context.Context, ch chan string){
	fmt.Println("run...")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}


// iota
func lesson70(){
//	fmt.Println(c1, c2, c3)
	fmt.Println(KOHE, KB, MB, GB)
}

const (
	KOHE = iota
	KB int = 1 << (10 * iota)
	MB
	GB 
)

// constが連番を使うときに使用する
const (
	c1 = iota
	c2 = iota
	c3
)


// sort
func lesson69(){
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool{return p[i].Age < p[j].Age })
		
	fmt.Println(i, s, p)
}

// regex
func lesson68(){
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
	
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)
	
	// s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println("fss=",fss," fss[0]=", fss[0], " fss[1]=", fss[1], " fss[2]=", fss[2])
}

// time
func lesson67(){
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), " ", t.Month(), " ", t.Day(), " ", t.Hour(), " ", t.Minute(), " ", t.Second())
}

func lesson64() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := alias.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}

// 20190422 Public and private
// 名称の先頭が大文字：public / 小文字:private
func lesson59(){
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mylib.Average(s))
	
	mylib.Say()
	
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
}