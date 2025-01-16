# Calculator_Go-lang
This is a simple calculator package written in Go, demonstrating basic mathematical operations with support for unit testing and benchmarking.

## Package Contents

- `calculator.go`: Contains the main mathematical functions.
- `calculator_test.go`: Contains the test cases and benchmarks for the package.

### Available Functions

1. **`Divide(a, b int) (int, error)`**
   - Performs integer division.
   - Returns an error if division by zero is attempted.

2. **`Square(a int) int`**
   - Returns the square of a number.

3. **`Add(a, b int) int`**
   - Returns the sum of two numbers.

4. **`Subtract(a, b int) int`**
   - Returns the difference between two numbers.

5. **`Multiply(a, b int) int`**
   - Returns the product of two numbers.

6. **`Factorial(n int) int`**
   - Returns the factorial of a number.

## Running Tests

To ensure your functions are working correctly, you can run the tests with the following commands:

### On Windows:

```powershell
cd \path\to\project
go test -v
```

### Testing in Go

- The `calculator_test.go` file contains the tests for all the functions. You can run `go test` to check that everything works as intended.

## Test Coverage

Test coverage refers to how much of your code is being tested. It's important to test various scenarios like normal cases, edge cases, and error handling to ensure your functions perform well in all situations.

### What Do We Test?

We ensure that:

- **Normal cases** (e.g., basic arithmetic like `2 ÷ 2`).
- **Edge cases** (e.g., division by zero).
- **Error handling** (ensuring errors are caught, such as trying to divide by zero).
- **Helper functions** like `Square` and `Factorial`.

### Checking Test Coverage

To check test coverage, run:

```bash
go test -cover
```

This will display how much of your code has been covered by tests, giving you an idea of how thorough your tests are.

## Benchmarking

Benchmarking helps you evaluate how fast your code runs and find areas for optimization.

### Why Benchmark?

- **Identify Performance Bottlenecks**: Find out which parts of your code are slower than expected.
- **Optimize Code**: Enhance performance by refining the slow parts of the code.
- **Compare Approaches**: Test different implementations and choose the most efficient.

### Running Benchmarks

To run the benchmarks on macOS or Windows, use the following command:

**Windows OS**

```powershell
go test -bench .
```

### Notes

- Files must be in the same directory.
- Test files must end with `_test.go`.
- Benchmark functions should start with `Benchmark`.

## Conclusion

The `Calculator Go` package provides basic mathematical operations such as division, square, addition, subtraction, multiplication, and factorial, along with tests and benchmarks. The tests ensure correctness, while the benchmarks help measure and optimize the performance of your functions.

