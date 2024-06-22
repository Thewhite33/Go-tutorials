package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	numbers := []int{1, 2, 3, 4}
	for index, val := range numbers {
		fmt.Printf("Index is %d and number is %d\n", index, val)
	}
}
