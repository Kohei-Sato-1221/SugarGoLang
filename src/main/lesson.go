package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	lesson17()
}

// Map
func lesson17(){
	m := map[string] int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(m["banana"])
	
	m["orange"] = 500
	fmt.Println(m)
	
	fmt.Println(m["nothing"]) //無いものを取り出そうとすると0になる！
	
	//Mapに含まれているかの確認
	v, ok := m["apple"]
	fmt.Println(v, ok) // 100 trueと表示
	
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2) // 0 falseと表示	
	
	m2 := make(map[string]int) //メモリ上に空のマップを作ってから要素の追加も可能
	m2["pc"] = 5000
	fmt.Println(m2)
	
	/* 以下はMapがNilなのでエラーになってしまう
	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m2)
	*/
	
	var s []int
	if s == nil{
		fmt.Println("s is Nil")
	}
}

// Slice
func lesson16(){
	n := make([]int, 3 , 5) 
	//長さが３、キャパシティが５のスライスを作る
	//長さ→初期値が０として要素が存在 キャパシティ→メモリが確保されるだけで要素はない状態
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	
	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	
	
	c = make([]int, 5) //要素数は5で、0が5つある状態
//	c = make([]int, 0, 5) //要素数は1で、0が1つ
	for i := 0; i < 5; i++{
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}

func lesson15(){
	// スライスは要素数を変更できる！
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])
	
	n[2] = 100
	fmt.Println(n)
	n = append(n, 100)
	fmt.Println(n)

	// ２次元配列のやり方
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	
	fmt.Println(board)
	
}

func lesson14(){
	var a [3] int
	a[0] = 100
	a[1] = 200
	a[2] = 300
	fmt.Println(a)
	
	/*
	var b [3]int = [3]int{100, 200, 300} //配列はサイズを変更できない！
	fmt.Println(b)
	*/
	
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}


func lesson13(){
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)
	
	var s string = "14"
//	z = int(s)	 この型変換はできない！
	i, err := strconv.Atoi(s)  //Atoi: アスキーtoインテジャー, 返り値が複数ある場合
	if err == nil{
		fmt.Println("## No error! ##")
	}
	fmt.Printf("%T %v", i, i)
}

func lesson12(){
//	var t, f bool = true, false
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, t) //%tとするとboolじゃないと表示できない！
	fmt.Printf("%T %v %t\n", f, f, t)
	
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}


func lesson11() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")
	fmt.Println(string("Hello World"[0])) //stringにキャストしないと72のアスキーコードが表示されてしまう
	
	var s string = "Hello World"
	// s[0] = "x" これはエラー
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s) // s自体は置き換わっていない！
	
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s) // sが変わっているはず
	
	fmt.Println(strings.Contains(s, "World"))
	
	fmt.Println(`改行テスト
	  \"
	改行テスト終わり`)
	
	fmt.Println("\" 文字列の中で\"を使いたい場合！")
}



// Lesson3 数値型
func lesson3Main(){
	// Goでは一番長い行に合わせてスペースを置くのが推奨されている
	/*
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8) //%Tは型, %vは値を表示
	*/
	
	/*
	x := 1 + 1
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)
	*/
	
	/* インクリメント
	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)
	*/
	
	/* シフト演算 */
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}


// Lesson2
const Pi = 3.14 //定数 constは宣言時には型を宣言しない！
const (
		Username = "test_user"
		Password = "test_pass"
)

//var big int = 9223372036854775807 + 1  overflowするパターン！
const big = 9223372036854775807 + 1

func lesson2Main(){
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1) //OverFlowしない！
}


// Lesson1
func lesson1Main(){
	/* 変数宣言の方法（Varの変数宣言は、関数外でも可能）
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t,f bool = true, false
	*/
	fmt.Println(i, f64, s, t, f)
	fmt.Println("----")
	foo()
}

var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t,f bool = true, false
)

func foo(){
	xi := 1 //ショートカット変数宣言
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}
