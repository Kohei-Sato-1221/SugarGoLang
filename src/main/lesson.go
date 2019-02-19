package main

import "fmt"

func main() {
	lesson3Main()
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
	fmt.Println(big - 1)
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
