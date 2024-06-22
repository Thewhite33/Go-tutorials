package main

import "fmt"

// func simpleFunc() {
// 	fmt.Println("Simple fxn")
// }
func addTwo(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello")
	//simpleFunc()
	ans := addTwo(2, 4)
	fmt.Println(ans)
}
