package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// println("Enter your name:")
	// fmt.Scan(&name)
	// fmt.Println("Name is", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Name is", name)
}
