package main

import "fmt"

/*
関数に関数を渡す式です
これはテスト記述なので、無視していいです。
*/

func first() func() {
	return func() {
		n := "Hello,World"
		fmt.Println(n)
	}
}

func main() {
	f := first()
	f()
}
