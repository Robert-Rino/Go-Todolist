package practices

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))

	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		result[i] = runes[j]
	}

	return string(result)
}

// RunReverseString demonstrates the reverseString function with interactive input
func RunReverseString() {
	fmt.Println("Waiting for input..., press Ctrl+D to exit")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(reverseString(line))
	}
}
