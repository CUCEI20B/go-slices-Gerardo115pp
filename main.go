package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), "\n")
	fmt.Println(inputs[0])
}
