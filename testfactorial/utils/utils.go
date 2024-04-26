package utils

import (
    "math/big"
)

// Factorial calculates the factorial of a given number using big.Int for large number computation.
func Factorial(n int) *big.Int {
    result := big.NewInt(1) // Initializes to 1
    if n > 0 {
        for i := 1; i <= n; i++ {
            result.Mul(result, big.NewInt(int64(i))) // Multiplies result by i
        }
    }
    return result
}
