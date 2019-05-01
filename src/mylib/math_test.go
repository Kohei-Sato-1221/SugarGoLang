package mylib

import (
	"testing"
	"fmt"
	"strconv"
)

var Debug bool = false

// lesson62 testing
func TestAverage(t *testing.T){
	if Debug {
		t.Skip("スキップした理由を記載し、終了")		
	}
	var v int
	v = Average([] int{1, 2, 3, 4, 5, 7})
	fmt.Println("平均は...." + strconv.Itoa(v))
	if v != 3{
		t.Error("Exepcted 3, got", v)
	}
}

func Example(){
	var v int
	v = Average([] int{1, 2, 3, 4, 5, 7})
	fmt.Println("平均は...." + strconv.Itoa(v))
}

func ExampleAverage(){
	var v int
	v = Average([] int{1, 2, 3, 4, 5, 7})
	fmt.Println("平均は...." + strconv.Itoa(v))
}

// terminalで実行する場合
// cd /Users/koheisato/eclipse-workspace/SugarGoLang/src
// go test ./... -count=1

// キャッシュを削除したい場合
// go clean -testcache
