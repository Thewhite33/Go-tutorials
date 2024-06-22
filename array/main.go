package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "Rohit"
	var price = [5]int{1, 2, 3}
	fmt.Println(name)
	fmt.Println(price[4])
	fmt.Println(len(price))
}
