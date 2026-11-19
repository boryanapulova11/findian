package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	s := strings.ToLower(strings.TrimSpace(input))

	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
