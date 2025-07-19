package practices

import (
	"testing"
)

// Tests for fibonacci function
func TestFibonacci(t *testing.T) {
	f := fibonacci()

	// Test the first 10 fibonacci numbers
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	for i, expectedValue := range expected {
		result := f()
		if result != expectedValue {
			t.Errorf("fibonacci() call %d: expected %d, got %d", i+1, expectedValue, result)
		}
	}
}

func TestFibonacciMultipleGenerators(t *testing.T) {
	// Test that multiple generators work independently
	f1 := fibonacci()
	f2 := fibonacci()

	// Both should start with 0
	if f1() != 0 {
		t.Error("First generator should start with 0")
	}
	if f2() != 0 {
		t.Error("Second generator should start with 0")
	}

	// Both should continue with 1
	if f1() != 1 {
		t.Error("First generator second call should be 1")
	}
	if f2() != 1 {
		t.Error("Second generator second call should be 1")
	}
}

// Tests for fizzbuzz function
func TestFizzbuzz(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// FizzBuzz cases (divisible by 15)
		{"15", "FizzBuzz"},
		{"30", "FizzBuzz"},
		{"45", "FizzBuzz"},
		{"0", "FizzBuzz"},

		// Fizz cases (divisible by 3 but not 5)
		{"3", "Fizz"},
		{"6", "Fizz"},
		{"9", "Fizz"},
		{"12", "Fizz"},

		// Buzz cases (divisible by 5 but not 3)
		{"5", "Buzz"},
		{"10", "Buzz"},
		{"20", "Buzz"},
		{"25", "Buzz"},

		// Regular numbers
		{"1", "1"},
		{"2", "2"},
		{"4", "4"},
		{"7", "7"},
		{"8", "8"},
		{"11", "11"},

		// Invalid input (non-numbers)
		{"abc", "abc"},
		{"", ""},
		{"hello", "hello"},
		{"12.5", "12.5"},

		// Negative numbers
		{"-15", "FizzBuzz"},
		{"-3", "Fizz"},
		{"-5", "Buzz"},
		{"-1", "-1"},
	}

	for _, tc := range testCases {
		result := fizzbuzz(tc.input)
		if result != tc.expected {
			t.Errorf("fizzbuzz(%q): expected %q, got %q", tc.input, tc.expected, result)
		}
	}
}

// Tests for isPalindrome function
func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		// True cases
		{"racecar", true},
		{"level", true},
		{"noon", true},
		{"a", true},
		{"", true},
		{"aba", true},
		{"abba", true},
		{"madam", true},

		// False cases
		{"hello", false},
		{"world", false},
		{"test", false},
		{"abc", false},
		{"ab", false},
		{"palindrome", false},

		// Numbers as strings
		{"121", true},
		{"12321", true},
		{"123", false},

		// Mixed case (current implementation is case-sensitive)
		{"Racecar", false}, // This will fail with current implementation
		{"RaceCar", false},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("isPalindrome(%q): expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

// Tests for reverseString function
func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"Go", "oG"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"racecar", "racecar"}, // palindrome
		{"12345", "54321"},
		{"Hello World!", "!dlroW olleH"},

		// Unicode characters
		{"café", "éfac"},
		{"🙂🙃", "🙃🙂"},
	}

	for _, tc := range testCases {
		result := reverseString(tc.input)
		if result != tc.expected {
			t.Errorf("reverseString(%q): expected %q, got %q", tc.input, tc.expected, result)
		}
	}
}

// Test that reverseString and isPalindrome work together correctly
func TestReverseStringWithPalindrome(t *testing.T) {
	testStrings := []string{"hello", "racecar", "level", "world"}

	for _, s := range testStrings {
		reversed := reverseString(s)
		isPalin := isPalindrome(s)

		// If original is palindrome, reverse should equal original
		if isPalin && reversed != s {
			t.Errorf("String %q is palindrome but reverse %q doesn't match", s, reversed)
		}

		// Double reverse should always equal original
		doubleReversed := reverseString(reversed)
		if doubleReversed != s {
			t.Errorf("Double reverse of %q should be %q, got %q", s, s, doubleReversed)
		}
	}
}

// Benchmark tests
func BenchmarkFibonacci(b *testing.B) {
	f := fibonacci()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkFizzbuzz(b *testing.B) {
	testInputs := []string{"1", "3", "5", "15", "abc"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range testInputs {
			fizzbuzz(input)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	testInputs := []string{"racecar", "hello", "level", "world", "a"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range testInputs {
			isPalindrome(input)
		}
	}
}

func BenchmarkReverseString(b *testing.B) {
	testInputs := []string{"hello", "world", "Go programming", "🙂🙃"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range testInputs {
			reverseString(input)
		}
	}
}
