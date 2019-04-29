package mylib

import (
	"fmt"
)

// Pが大文字なので、外から呼び出し　→　pにすると小文字
type Person struct{
	Name string
	Age int
}

func Say(){
	fmt.Println("hello! world!!")
}
