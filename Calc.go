package main

import (
	"errors"
	"fmt"
)

// Basic Operations
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Advanced Operations
func square(a int) int {
	return a * a
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Example usage
	fmt.Println("Addition:", Add(10, 5))
	fmt.Println("Subtraction:", Subtract(10, 5))
	fmt.Println("Multiplication:", Multiply(10, 5))

	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	fmt.Println("Square:", square(4))
	fmt.Println("Factorial:", Factorial(5))
}
