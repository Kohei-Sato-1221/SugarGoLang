package main2

import (
	"fmt"
	"main2/mylib" // 他ファイルのGOファイルを呼ぶ方法
)

func main(){
	lesson59()
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