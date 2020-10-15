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
	var output int = 0
	for h := 1; h < len(inputs); h++ {
		numericValue, _ := strconv.Atoi(inputs[h])
		output += numericValue
	}
	fmt.Println(output)
}
