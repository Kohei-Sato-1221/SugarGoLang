package mylib

import (
	"testing"
)

// lesson62 testing
fucn TestAverage(t *testing.T){
	v := Average([] int{1, 2, 3, 4, 5, 6})
	fmt.Println("Average is " + v)
	if v != 3{
		t.Error("Exepcted 3, got", v)
	}
}

// terminalで実行する場合
// ※プロジェクトのあるフォルダにターミナルで移動
// go test ./...