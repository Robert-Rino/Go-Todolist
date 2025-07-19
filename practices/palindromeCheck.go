package practices

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for {
		if left > right {
			return true
		}

		if s[left] != s[right] {
			return false
		}

		left += 1
		right -= 1
	}
}

// RunPalindromeCheck demonstrates the isPalindrome function with interactive input
func RunPalindromeCheck() {
	fmt.Println("Waiting for input..., press Ctrl+D to exit")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(isPalindrome(line))
	}
}
