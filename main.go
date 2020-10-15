package main

import (
	"fmt"
)

func main() {
	var lines int
	var holder int
	var output int
	fmt.Scanln(&lines)
	for h := 0; h < lines; h++ {
		fmt.Scanln(&holder)
		output += holder
	}
	fmt.Println(output)
}
