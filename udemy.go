package main

import "fmt"

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
