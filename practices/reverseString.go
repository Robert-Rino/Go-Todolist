package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)
	result := make([]rune, len(s))

	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		result[i] = runes[j]
	}

	return string(result)

}

func main() {
	fmt.Println("Waiting for input..., press Ctrl+D to exit")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(reverseString(line))
	}
}
