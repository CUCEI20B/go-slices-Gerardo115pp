package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	inputs := strings.Split(s, `\r\n`)
	if len(inputs) <= 1 {
		inputs = strings.Split(s, `\n`)
	}
	fmt.Println(inputs)
	var output int = 0
	for h := 1; h < len(inputs); h++ {
		numericValue, _ := strconv.Atoi(inputs[h])
		output += numericValue
	}
	fmt.Printf("%d", output)
}
