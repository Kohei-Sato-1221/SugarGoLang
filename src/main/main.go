package main

import (
	"github.com/markcheno/go-quote"
	alias "github.com/markcheno/go-talib"
	"fmt"
	"mylib" // 他ファイルのGOファイルを呼ぶ方法
)

func main(){
	lesson64()
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