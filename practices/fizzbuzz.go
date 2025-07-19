package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(s string) string {
	num, err := strconv.Atoi(s)

	if err != nil {
		return s
	}

	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return s
	}

}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Println("Waiting for input..., press Ctrl+D to exit")
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println("line:", line)

		fmt.Println(fizzbuzz(line))
	}

	if err := stdin.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
