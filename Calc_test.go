package main

import "testing"

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(3, 2)
	if result != 5 {
		t.Errorf("Add(3, 2) failed. Expected 5, got %v", result)
	}
}

// TestSubtract tests the Subtract function
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract(5, 3) failed. Expected 2, got %v", result)
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)
	if result != 12 {
		t.Errorf("Multiply(4, 3) failed. Expected 12, got %v", result)
	}
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	result, err := Divide(6, 2)
	if result != 3 || err != nil {
		t.Errorf("Divide(6, 2) failed. Expected 3, got %v", result)
	}

	result, err = Divide(6, 0)
	if err == nil {
		t.Errorf("Divide(6, 0) should fail but did not")
	}
}

// TestSquare tests the Square function
func TestSquare(t *testing.T) {
	result := square(4)
	if result != 16 {
		t.Errorf("Square(4) failed. Expected 16, got %v", result)
	}
}

// TestFactorial tests the Factorial function
func TestFactorial(t *testing.T) {
	result := Factorial(5)
	if result != 120 {
		t.Errorf("Factorial(5) failed. Expected 120, got %v", result)
	}
}

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3, 2)
	}
}

// BenchmarkSubtract benchmarks the Subtract function
func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(5, 3)
	}
}

// BenchmarkMultiply benchmarks the Multiply function
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(4, 3)
	}
}

// BenchmarkDivide benchmarks the Divide function
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(6, 2)
	}
}

// BenchmarkSquare benchmarks the Square function
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		square(4)
	}
}

// BenchmarkFactorial benchmarks the Factorial function
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(5)
	}
}
