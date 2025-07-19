# Go Programming Practices

This repository contains various Go programming practice exercises and their implementations.

## 📁 Practice Programs

### Available Programs:

1. **Fibonacci Generator** (`fibonacci.go`)
   - Generates fibonacci sequence using closures
   - Run: `go run practices/fibonacci.go`

2. **FizzBuzz** (`fizzbuzz.go`)
   - Classic FizzBuzz implementation with input handling
   - Run: `go run practices/fizzbuzz.go`

3. **Palindrome Checker** (`palindromeCheck.go`)
   - Checks if a string is a palindrome
   - Run: `go run practices/palindromeCheck.go`

4. **String Reverser** (`reverseString.go`)
   - Reverses strings with Unicode support
   - Run: `go run practices/reverseString.go`

## 🧪 Testing

### Test File
All functions are comprehensively tested in a single test file: `practices/practices_test.go`

### Running Tests

#### Basic Test Commands:
```bash
# Navigate to practices directory
cd practices

# Run all tests
go test

# Run tests with verbose output (shows individual test results)
go test -v

# Run tests with coverage report
go test -cover
```

#### Running Specific Tests:
```bash
# Test only fibonacci function
go test -run TestFibonacci

# Test only fizzbuzz function
go test -run TestFizzbuzz

# Test only palindrome function
go test -run TestIsPalindrome

# Test only reverse string function
go test -run TestReverseString
```

#### Performance Benchmarks:
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFibonacci
go test -bench=BenchmarkFizzbuzz
go test -bench=BenchmarkIsPalindrome
go test -bench=BenchmarkReverseString
```

#### Advanced Testing:
```bash
# Generate detailed coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run tests multiple times
go test -count=3

# Run tests with race condition detection
go test -race
```

### Test Coverage

The test file includes comprehensive coverage for:

- ✅ **Fibonacci**: Sequence accuracy, multiple generators, edge cases
- ✅ **FizzBuzz**: All divisibility rules, invalid input, negative numbers
- ✅ **Palindrome**: Various palindromes, non-palindromes, edge cases
- ✅ **Reverse String**: Basic reversal, Unicode characters, empty strings
- ✅ **Integration**: Tests how functions work together
- ✅ **Performance**: Benchmark tests for all functions

